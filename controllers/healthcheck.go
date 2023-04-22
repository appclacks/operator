package controllers

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	appclackscomv1alpha1 "appclacks.com/operator/api/v1alpha1"
	goclient "github.com/appclacks/cli/client"
	apitypes "github.com/appclacks/go-types"
	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

const (
	defaultTimeout = 10 * time.Second

	created       = "Created"
	finalizerName = "healthchecks.appclacks.com/finalizer"
)

func createHealthcheck(ctx context.Context, client client.Client, appclacks *goclient.Client, check client.Object) error {
	switch healthcheck := check.(type) {
	case *appclackscomv1alpha1.CommandHealthcheck:
		payload := apitypes.CreateCommandHealthcheckInput{
			Name:        healthcheck.Name,
			Description: healthcheck.Spec.Description,
			Labels:      healthcheck.Labels,
			Timeout:     healthcheck.Spec.Timeout,
			Interval:    healthcheck.Spec.Interval,
			Enabled:     healthcheck.Spec.Enabled,
			HealthcheckCommandDefinition: apitypes.HealthcheckCommandDefinition{
				Command:   healthcheck.Spec.Command,
				Arguments: healthcheck.Spec.Arguments,
			},
		}

		ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
		defer cancel()
		_, err := appclacks.CreateCommandHealthcheck(ctx, payload)
		if err != nil {
			return fmt.Errorf("fail to create healthcheck: %w", err)
		}
		healthcheck.Status.State = created
	case *appclackscomv1alpha1.HTTPHealthcheck:
		payload := apitypes.CreateHTTPHealthcheckInput{
			Name:        healthcheck.Name,
			Description: healthcheck.Spec.Description,
			Labels:      healthcheck.Labels,
			Timeout:     healthcheck.Spec.Timeout,
			Interval:    healthcheck.Spec.Interval,
			Enabled:     healthcheck.Spec.Enabled,
			HealthcheckHTTPDefinition: apitypes.HealthcheckHTTPDefinition{
				ValidStatus: healthcheck.Spec.ValidStatus,
				Target:      healthcheck.Spec.Target,
				Method:      healthcheck.Spec.Method,
				Port:        healthcheck.Spec.Port,
				Protocol:    healthcheck.Spec.Protocol,
				Redirect:    healthcheck.Spec.Redirect,
				Query:       healthcheck.Spec.Query,
				Body:        healthcheck.Spec.Body,
				BodyRegexp:  healthcheck.Spec.BodyRegexp,
				Headers:     healthcheck.Spec.Headers,
				Path:        healthcheck.Spec.Path,
				Key:         healthcheck.Spec.Key,
				Cert:        healthcheck.Spec.Cert,
				Cacert:      healthcheck.Spec.Cacert,
				Insecure:    healthcheck.Spec.Insecure,
				ServerName:  healthcheck.Spec.ServerName,
				Host:        healthcheck.Spec.Host,
			},
		}

		ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
		defer cancel()
		_, err := appclacks.CreateHTTPHealthcheck(ctx, payload)
		if err != nil {
			return fmt.Errorf("fail to create healthcheck: %w", err)
		}
		healthcheck.Status.State = created
	case *appclackscomv1alpha1.TLSHealthcheck:
		payload := apitypes.CreateTLSHealthcheckInput{
			Name:        healthcheck.Name,
			Description: healthcheck.Spec.Description,
			Labels:      healthcheck.Labels,
			Timeout:     healthcheck.Spec.Timeout,
			Interval:    healthcheck.Spec.Interval,
			Enabled:     healthcheck.Spec.Enabled,
			HealthcheckTLSDefinition: apitypes.HealthcheckTLSDefinition{
				Target:          healthcheck.Spec.Target,
				Port:            healthcheck.Spec.Port,
				Key:             healthcheck.Spec.Key,
				Cert:            healthcheck.Spec.Cert,
				Cacert:          healthcheck.Spec.Cacert,
				ServerName:      healthcheck.Spec.ServerName,
				Insecure:        healthcheck.Spec.Insecure,
				ExpirationDelay: healthcheck.Spec.ExpirationDelay,
			},
		}

		ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
		defer cancel()
		_, err := appclacks.CreateTLSHealthcheck(ctx, payload)
		if err != nil {
			return fmt.Errorf("fail to create healthcheck: %w", err)
		}
		healthcheck.Status.State = created
	case *appclackscomv1alpha1.TCPHealthcheck:
		payload := apitypes.CreateTCPHealthcheckInput{
			Name:        healthcheck.Name,
			Description: healthcheck.Spec.Description,
			Labels:      healthcheck.Labels,
			Timeout:     healthcheck.Spec.Timeout,
			Interval:    healthcheck.Spec.Interval,
			Enabled:     healthcheck.Spec.Enabled,
			HealthcheckTCPDefinition: apitypes.HealthcheckTCPDefinition{
				Target:     healthcheck.Spec.Target,
				Port:       healthcheck.Spec.Port,
				ShouldFail: healthcheck.Spec.ShouldFail,
			},
		}

		ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
		defer cancel()
		_, err := appclacks.CreateTCPHealthcheck(ctx, payload)
		if err != nil {
			return fmt.Errorf("fail to create healthcheck: %w", err)
		}
		healthcheck.Status.State = created
	case *appclackscomv1alpha1.DNSHealthcheck:
		payload := apitypes.CreateDNSHealthcheckInput{
			Name:        healthcheck.Name,
			Description: healthcheck.Spec.Description,
			Labels:      healthcheck.Labels,
			Timeout:     healthcheck.Spec.Timeout,
			Interval:    healthcheck.Spec.Interval,
			Enabled:     healthcheck.Spec.Enabled,
			HealthcheckDNSDefinition: apitypes.HealthcheckDNSDefinition{
				Domain:      healthcheck.Spec.Domain,
				ExpectedIPs: healthcheck.Spec.ExpectedIPs,
			},
		}

		ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
		defer cancel()
		_, err := appclacks.CreateDNSHealthcheck(ctx, payload)
		if err != nil {
			return fmt.Errorf("fail to create healthcheck: %w", err)
		}
		healthcheck.Status.State = created
	default:
		return fmt.Errorf("invalid healthcheck type %s", reflect.TypeOf(check))
	}
	if !controllerutil.ContainsFinalizer(check, finalizerName) {
		controllerutil.AddFinalizer(check, finalizerName)
	}
	if err := client.Update(ctx, check); err != nil {
		return fmt.Errorf("fail to update healthcheck in Kubernetes: %w", err)
	}
	return nil
}

func updateHealthcheck(ctx context.Context, client client.Client, appclacks *goclient.Client, check client.Object, existingCheck apitypes.Healthcheck) error {
	// the healthcheck does not contain finalizers: should be added
	k8sUpdate := false
	if !controllerutil.ContainsFinalizer(check, finalizerName) {
		k8sUpdate = true
		controllerutil.AddFinalizer(check, finalizerName)
	}

	switch healthcheck := check.(type) {
	case *appclackscomv1alpha1.CommandHealthcheck:
		definition, ok := existingCheck.Definition.(apitypes.HealthcheckCommandDefinition)
		if !ok {
			return errors.New("invalid dns healthcheck definition")
		}
		if healthcheck.Spec.Description != existingCheck.Description ||
			!reflect.DeepEqual(healthcheck.Labels, existingCheck.Labels) ||
			healthcheck.Spec.Timeout != existingCheck.Timeout ||
			healthcheck.Spec.Interval != existingCheck.Interval ||
			healthcheck.Spec.Enabled != existingCheck.Enabled ||
			healthcheck.Spec.Command != definition.Command ||
			!reflect.DeepEqual(healthcheck.Spec.Arguments, definition.Arguments) {
			payload := apitypes.UpdateCommandHealthcheckInput{
				ID:          existingCheck.ID,
				Name:        healthcheck.Name,
				Description: healthcheck.Spec.Description,
				Labels:      healthcheck.Labels,
				Timeout:     healthcheck.Spec.Timeout,
				Interval:    healthcheck.Spec.Interval,
				Enabled:     healthcheck.Spec.Enabled,
				HealthcheckCommandDefinition: apitypes.HealthcheckCommandDefinition{
					Command:   healthcheck.Spec.Command,
					Arguments: healthcheck.Spec.Arguments,
				},
			}
			ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
			defer cancel()
			_, err := appclacks.UpdateCommandHealthcheck(ctx, payload)
			if err != nil {
				return fmt.Errorf("fail to update healthcheck: %w", err)
			}
			healthcheck.Status.State = created
		}
	case *appclackscomv1alpha1.HTTPHealthcheck:
		definition, ok := existingCheck.Definition.(apitypes.HealthcheckHTTPDefinition)
		if !ok {
			return errors.New("invalid HTTP healthcheck definition")
		}
		if healthcheck.Spec.Description != existingCheck.Description ||
			!reflect.DeepEqual(healthcheck.Labels, existingCheck.Labels) ||
			healthcheck.Spec.Timeout != existingCheck.Timeout ||
			healthcheck.Spec.Interval != existingCheck.Interval ||
			healthcheck.Spec.Enabled != existingCheck.Enabled ||
			!reflect.DeepEqual(healthcheck.Spec.ValidStatus, definition.ValidStatus) ||
			healthcheck.Spec.Target != definition.Target ||
			healthcheck.Spec.Method != definition.Method ||
			healthcheck.Spec.Port != definition.Port ||
			healthcheck.Spec.Protocol != definition.Protocol ||
			healthcheck.Spec.Redirect != definition.Redirect ||
			!reflect.DeepEqual(healthcheck.Spec.Query, definition.Query) ||
			healthcheck.Spec.Body != definition.Body ||
			!reflect.DeepEqual(healthcheck.Spec.BodyRegexp, definition.BodyRegexp) ||
			!reflect.DeepEqual(healthcheck.Spec.Headers, definition.Headers) ||
			healthcheck.Spec.Path != definition.Path ||
			healthcheck.Spec.Key != definition.Key ||
			healthcheck.Spec.Cert != definition.Cert ||
			healthcheck.Spec.Cacert != definition.Cacert ||
			healthcheck.Spec.Insecure != definition.Insecure ||
			healthcheck.Spec.ServerName != definition.ServerName ||
			healthcheck.Spec.Host != definition.Host {
			payload := apitypes.UpdateHTTPHealthcheckInput{
				ID:          existingCheck.ID,
				Name:        healthcheck.Name,
				Description: healthcheck.Spec.Description,
				Labels:      healthcheck.Labels,
				Timeout:     healthcheck.Spec.Timeout,
				Interval:    healthcheck.Spec.Interval,
				Enabled:     healthcheck.Spec.Enabled,
				HealthcheckHTTPDefinition: apitypes.HealthcheckHTTPDefinition{
					ValidStatus: healthcheck.Spec.ValidStatus,
					Target:      healthcheck.Spec.Target,
					Method:      healthcheck.Spec.Method,
					Port:        healthcheck.Spec.Port,
					Protocol:    healthcheck.Spec.Protocol,
					Redirect:    healthcheck.Spec.Redirect,
					Query:       healthcheck.Spec.Query,
					Body:        healthcheck.Spec.Body,
					BodyRegexp:  healthcheck.Spec.BodyRegexp,
					Headers:     healthcheck.Spec.Headers,
					Path:        healthcheck.Spec.Path,
					Key:         healthcheck.Spec.Key,
					Cert:        healthcheck.Spec.Cert,
					Cacert:      healthcheck.Spec.Cacert,
					Insecure:    healthcheck.Spec.Insecure,
					ServerName:  healthcheck.Spec.ServerName,
					Host:        healthcheck.Spec.Host,
				},
			}
			ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
			defer cancel()
			_, err := appclacks.UpdateHTTPHealthcheck(ctx, payload)
			if err != nil {
				return fmt.Errorf("fail to update healthcheck: %w", err)
			}
			healthcheck.Status.State = created
		}
	case *appclackscomv1alpha1.TLSHealthcheck:
		definition, ok := existingCheck.Definition.(apitypes.HealthcheckTLSDefinition)
		if !ok {
			return errors.New("invalid TLS healthcheck definition")
		}
		if healthcheck.Spec.Description != existingCheck.Description ||
			!reflect.DeepEqual(healthcheck.Labels, existingCheck.Labels) ||
			healthcheck.Spec.Timeout != existingCheck.Timeout ||
			healthcheck.Spec.Interval != existingCheck.Interval ||
			healthcheck.Spec.Enabled != existingCheck.Enabled ||
			healthcheck.Spec.Target != definition.Target ||
			healthcheck.Spec.Port != definition.Port ||
			healthcheck.Spec.Key != definition.Key ||
			healthcheck.Spec.Cert != definition.Cert ||
			healthcheck.Spec.Cacert != definition.Cacert ||
			healthcheck.Spec.ServerName != definition.ServerName ||
			healthcheck.Spec.Insecure != definition.Insecure ||
			healthcheck.Spec.ExpirationDelay != definition.ExpirationDelay {
			payload := apitypes.UpdateTLSHealthcheckInput{
				ID:          existingCheck.ID,
				Name:        healthcheck.Name,
				Description: healthcheck.Spec.Description,
				Labels:      healthcheck.Labels,
				Timeout:     healthcheck.Spec.Timeout,
				Interval:    healthcheck.Spec.Interval,
				Enabled:     healthcheck.Spec.Enabled,
				HealthcheckTLSDefinition: apitypes.HealthcheckTLSDefinition{
					Target:          healthcheck.Spec.Target,
					Port:            healthcheck.Spec.Port,
					Key:             healthcheck.Spec.Key,
					Cert:            healthcheck.Spec.Cert,
					Cacert:          healthcheck.Spec.Cacert,
					ServerName:      healthcheck.Spec.ServerName,
					Insecure:        healthcheck.Spec.Insecure,
					ExpirationDelay: healthcheck.Spec.ExpirationDelay,
				},
			}
			ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
			defer cancel()
			_, err := appclacks.UpdateTLSHealthcheck(ctx, payload)
			if err != nil {
				return fmt.Errorf("fail to update healthcheck: %w", err)
			}
			healthcheck.Status.State = created
		}
	case *appclackscomv1alpha1.TCPHealthcheck:
		definition, ok := existingCheck.Definition.(apitypes.HealthcheckTCPDefinition)
		if !ok {
			return errors.New("invalid TCP healthcheck definition")
		}
		if healthcheck.Spec.Description != existingCheck.Description ||
			!reflect.DeepEqual(healthcheck.Labels, existingCheck.Labels) ||
			healthcheck.Spec.Timeout != existingCheck.Timeout ||
			healthcheck.Spec.Interval != existingCheck.Interval ||
			healthcheck.Spec.Enabled != existingCheck.Enabled ||
			healthcheck.Spec.Target != definition.Target ||
			healthcheck.Spec.Port != definition.Port ||
			healthcheck.Spec.ShouldFail != definition.ShouldFail {
			payload := apitypes.UpdateTCPHealthcheckInput{
				ID:          existingCheck.ID,
				Name:        healthcheck.Name,
				Description: healthcheck.Spec.Description,
				Labels:      healthcheck.Labels,
				Timeout:     healthcheck.Spec.Timeout,
				Interval:    healthcheck.Spec.Interval,
				Enabled:     healthcheck.Spec.Enabled,
				HealthcheckTCPDefinition: apitypes.HealthcheckTCPDefinition{
					Target:     healthcheck.Spec.Target,
					Port:       healthcheck.Spec.Port,
					ShouldFail: healthcheck.Spec.ShouldFail,
				},
			}
			ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
			defer cancel()
			_, err := appclacks.UpdateTCPHealthcheck(ctx, payload)
			if err != nil {
				return fmt.Errorf("fail to update healthcheck: %w", err)
			}
			healthcheck.Status.State = created
		}
	case *appclackscomv1alpha1.DNSHealthcheck:
		definition, ok := existingCheck.Definition.(apitypes.HealthcheckDNSDefinition)
		if !ok {
			return errors.New("invalid dns healthcheck definition")
		}
		if healthcheck.Spec.Description != existingCheck.Description ||
			!reflect.DeepEqual(healthcheck.Labels, existingCheck.Labels) ||
			healthcheck.Spec.Timeout != existingCheck.Timeout ||
			healthcheck.Spec.Interval != existingCheck.Interval ||
			healthcheck.Spec.Enabled != existingCheck.Enabled ||
			healthcheck.Spec.Domain != definition.Domain ||
			!reflect.DeepEqual(healthcheck.Spec.ExpectedIPs, definition.ExpectedIPs) {
			payload := apitypes.UpdateDNSHealthcheckInput{
				ID:          existingCheck.ID,
				Name:        healthcheck.Name,
				Description: healthcheck.Spec.Description,
				Labels:      healthcheck.Labels,
				Timeout:     healthcheck.Spec.Timeout,
				Interval:    healthcheck.Spec.Interval,
				Enabled:     healthcheck.Spec.Enabled,
				HealthcheckDNSDefinition: apitypes.HealthcheckDNSDefinition{
					Domain:      healthcheck.Spec.Domain,
					ExpectedIPs: healthcheck.Spec.ExpectedIPs,
				},
			}
			ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
			defer cancel()
			_, err := appclacks.UpdateDNSHealthcheck(ctx, payload)
			if err != nil {
				return fmt.Errorf("fail to update healthcheck: %w", err)
			}
			healthcheck.Status.State = created
		}
	default:
		return fmt.Errorf("invalid healthcheck type %s", reflect.TypeOf(check))
	}
	if k8sUpdate {
		if err := client.Update(ctx, check); err != nil {
			return fmt.Errorf("fail to update healthcheck in Kubernetes: %w", err)
		}
	}
	return nil
}

func deleteHealthcheck(ctx context.Context, client client.Client, appclacks *goclient.Client, check client.Object, id string, exists bool) error {
	if controllerutil.ContainsFinalizer(check, finalizerName) {
		if exists {
			ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
			defer cancel()
			_, err := appclacks.DeleteHealthcheck(ctx, apitypes.DeleteHealthcheckInput{
				ID: id,
			})

			if err != nil {
				return fmt.Errorf("fail to delete healthcheck from Appclacks: %w", err)
			}
		}

		controllerutil.RemoveFinalizer(check, finalizerName)
		if err := client.Update(ctx, check); err != nil {
			return err
		}
	}
	return nil
}

func reconcile(ctx context.Context, log logr.Logger, client client.Client, appclacks *goclient.Client, check client.Object, name string) error {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	existingCheck, getError := appclacks.GetHealthcheck(ctx, apitypes.GetHealthcheckInput{
		Identifier: name,
	})

	if getError != nil && getError != goclient.ErrNotFound {
		return getError
	}

	if !check.GetDeletionTimestamp().IsZero() {
		err := deleteHealthcheck(ctx, client, appclacks, check, existingCheck.ID, getError == nil)
		if err != nil {
			return err
		}
		log.Info("healthcheck deleted")
		return nil
	}

	if getError != nil && getError == goclient.ErrNotFound {
		err := createHealthcheck(ctx, client, appclacks, check)
		if err != nil {
			return fmt.Errorf("fail to create healthcheck: %w", err)
		}

		log.Info("healthcheck created")
		return nil
	}

	err := updateHealthcheck(ctx, client, appclacks, check, existingCheck)
	if err != nil {
		return err
	}
	log.Info("healthcheck reconciled")
	return nil
}

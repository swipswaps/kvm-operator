package pvc

import (
	"context"
	"fmt"

	"github.com/giantswarm/microerror"
	corev1 "k8s.io/api/core/v1"

	"github.com/giantswarm/kvm-operator/service/controller/key"
)

func (r *Resource) GetDesiredState(ctx context.Context, obj interface{}) (interface{}, error) {
	customObject, err := key.ToCustomObject(obj)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	var PVCs []*corev1.PersistentVolumeClaim

	if key.StorageType(customObject) == "persistentVolume" {
		r.logger.LogCtx(ctx, "level", "debug", "message", "computing the new PVCs")

		PVCs, err = newEtcdPVCs(customObject)
		if err != nil {
			return nil, microerror.Mask(err)
		}

		r.logger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("computed the %d new PVCs", len(PVCs)))
	} else {
		r.logger.LogCtx(ctx, "level", "debug", "message", "not computing the new PVCs because storage type is not 'persistentVolume'")
	}

	return PVCs, nil
}

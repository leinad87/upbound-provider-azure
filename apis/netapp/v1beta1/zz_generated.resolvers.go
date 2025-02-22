/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-azure/apis/azure/v1beta1"
	v1beta11 "github.com/upbound/provider-azure/apis/network/v1beta1"
	rconfig "github.com/upbound/provider-azure/apis/rconfig"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Account.
func (mg *Account) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Pool.
func (mg *Pool) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AccountName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.AccountNameRef,
		Selector:     mg.Spec.ForProvider.AccountNameSelector,
		To: reference.To{
			List:    &AccountList{},
			Managed: &Account{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AccountName")
	}
	mg.Spec.ForProvider.AccountName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AccountNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Snapshot.
func (mg *Snapshot) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AccountName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.AccountNameRef,
		Selector:     mg.Spec.ForProvider.AccountNameSelector,
		To: reference.To{
			List:    &AccountList{},
			Managed: &Account{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AccountName")
	}
	mg.Spec.ForProvider.AccountName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AccountNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PoolName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.PoolNameRef,
		Selector:     mg.Spec.ForProvider.PoolNameSelector,
		To: reference.To{
			List:    &PoolList{},
			Managed: &Pool{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PoolName")
	}
	mg.Spec.ForProvider.PoolName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PoolNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VolumeName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VolumeNameRef,
		Selector:     mg.Spec.ForProvider.VolumeNameSelector,
		To: reference.To{
			List:    &VolumeList{},
			Managed: &Volume{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VolumeName")
	}
	mg.Spec.ForProvider.VolumeName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VolumeNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SnapshotPolicy.
func (mg *SnapshotPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AccountName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.AccountNameRef,
		Selector:     mg.Spec.ForProvider.AccountNameSelector,
		To: reference.To{
			List:    &AccountList{},
			Managed: &Account{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AccountName")
	}
	mg.Spec.ForProvider.AccountName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AccountNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Volume.
func (mg *Volume) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AccountName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.AccountNameRef,
		Selector:     mg.Spec.ForProvider.AccountNameSelector,
		To: reference.To{
			List:    &AccountList{},
			Managed: &Account{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AccountName")
	}
	mg.Spec.ForProvider.AccountName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AccountNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CreateFromSnapshotResourceID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.CreateFromSnapshotResourceIDRef,
		Selector:     mg.Spec.ForProvider.CreateFromSnapshotResourceIDSelector,
		To: reference.To{
			List:    &SnapshotList{},
			Managed: &Snapshot{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CreateFromSnapshotResourceID")
	}
	mg.Spec.ForProvider.CreateFromSnapshotResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CreateFromSnapshotResourceIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.DataProtectionReplication); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DataProtectionReplication[i3].RemoteVolumeResourceID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.DataProtectionReplication[i3].RemoteVolumeResourceIDRef,
			Selector:     mg.Spec.ForProvider.DataProtectionReplication[i3].RemoteVolumeResourceIDSelector,
			To: reference.To{
				List:    &VolumeList{},
				Managed: &Volume{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.DataProtectionReplication[i3].RemoteVolumeResourceID")
		}
		mg.Spec.ForProvider.DataProtectionReplication[i3].RemoteVolumeResourceID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.DataProtectionReplication[i3].RemoteVolumeResourceIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.DataProtectionSnapshotPolicy); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DataProtectionSnapshotPolicy[i3].SnapshotPolicyID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.DataProtectionSnapshotPolicy[i3].SnapshotPolicyIDRef,
			Selector:     mg.Spec.ForProvider.DataProtectionSnapshotPolicy[i3].SnapshotPolicyIDSelector,
			To: reference.To{
				List:    &SnapshotPolicyList{},
				Managed: &SnapshotPolicy{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.DataProtectionSnapshotPolicy[i3].SnapshotPolicyID")
		}
		mg.Spec.ForProvider.DataProtectionSnapshotPolicy[i3].SnapshotPolicyID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.DataProtectionSnapshotPolicy[i3].SnapshotPolicyIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PoolName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.PoolNameRef,
		Selector:     mg.Spec.ForProvider.PoolNameSelector,
		To: reference.To{
			List:    &PoolList{},
			Managed: &Pool{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PoolName")
	}
	mg.Spec.ForProvider.PoolName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PoolNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SubnetIDRef,
		Selector:     mg.Spec.ForProvider.SubnetIDSelector,
		To: reference.To{
			List:    &v1beta11.SubnetList{},
			Managed: &v1beta11.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetID")
	}
	mg.Spec.ForProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CreateFromSnapshotResourceID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.CreateFromSnapshotResourceIDRef,
		Selector:     mg.Spec.InitProvider.CreateFromSnapshotResourceIDSelector,
		To: reference.To{
			List:    &SnapshotList{},
			Managed: &Snapshot{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.CreateFromSnapshotResourceID")
	}
	mg.Spec.InitProvider.CreateFromSnapshotResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CreateFromSnapshotResourceIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.DataProtectionReplication); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DataProtectionReplication[i3].RemoteVolumeResourceID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.DataProtectionReplication[i3].RemoteVolumeResourceIDRef,
			Selector:     mg.Spec.InitProvider.DataProtectionReplication[i3].RemoteVolumeResourceIDSelector,
			To: reference.To{
				List:    &VolumeList{},
				Managed: &Volume{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.DataProtectionReplication[i3].RemoteVolumeResourceID")
		}
		mg.Spec.InitProvider.DataProtectionReplication[i3].RemoteVolumeResourceID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.DataProtectionReplication[i3].RemoteVolumeResourceIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.DataProtectionSnapshotPolicy); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DataProtectionSnapshotPolicy[i3].SnapshotPolicyID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.DataProtectionSnapshotPolicy[i3].SnapshotPolicyIDRef,
			Selector:     mg.Spec.InitProvider.DataProtectionSnapshotPolicy[i3].SnapshotPolicyIDSelector,
			To: reference.To{
				List:    &SnapshotPolicyList{},
				Managed: &SnapshotPolicy{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.DataProtectionSnapshotPolicy[i3].SnapshotPolicyID")
		}
		mg.Spec.InitProvider.DataProtectionSnapshotPolicy[i3].SnapshotPolicyID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.DataProtectionSnapshotPolicy[i3].SnapshotPolicyIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SubnetID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.SubnetIDRef,
		Selector:     mg.Spec.InitProvider.SubnetIDSelector,
		To: reference.To{
			List:    &v1beta11.SubnetList{},
			Managed: &v1beta11.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SubnetID")
	}
	mg.Spec.InitProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SubnetIDRef = rsp.ResolvedReference

	return nil
}

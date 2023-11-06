//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Bucket) DeepCopyInto(out *S3Bucket) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Bucket.
func (in *S3Bucket) DeepCopy() *S3Bucket {
	if in == nil {
		return nil
	}
	out := new(S3Bucket)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S3Bucket) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3BucketList) DeepCopyInto(out *S3BucketList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]S3Bucket, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3BucketList.
func (in *S3BucketList) DeepCopy() *S3BucketList {
	if in == nil {
		return nil
	}
	out := new(S3BucketList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S3BucketList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3BucketSpec) DeepCopyInto(out *S3BucketSpec) {
	*out = *in
	if in.S3SubuserBinding != nil {
		in, out := &in.S3SubuserBinding, &out.S3SubuserBinding
		*out = make([]SubuserBinding, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3BucketSpec.
func (in *S3BucketSpec) DeepCopy() *S3BucketSpec {
	if in == nil {
		return nil
	}
	out := new(S3BucketSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3BucketStatus) DeepCopyInto(out *S3BucketStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3BucketStatus.
func (in *S3BucketStatus) DeepCopy() *S3BucketStatus {
	if in == nil {
		return nil
	}
	out := new(S3BucketStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3User) DeepCopyInto(out *S3User) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3User.
func (in *S3User) DeepCopy() *S3User {
	if in == nil {
		return nil
	}
	out := new(S3User)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S3User) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3UserClaim) DeepCopyInto(out *S3UserClaim) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3UserClaim.
func (in *S3UserClaim) DeepCopy() *S3UserClaim {
	if in == nil {
		return nil
	}
	out := new(S3UserClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S3UserClaim) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3UserClaimList) DeepCopyInto(out *S3UserClaimList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]S3UserClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3UserClaimList.
func (in *S3UserClaimList) DeepCopy() *S3UserClaimList {
	if in == nil {
		return nil
	}
	out := new(S3UserClaimList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S3UserClaimList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3UserClaimSpec) DeepCopyInto(out *S3UserClaimSpec) {
	*out = *in
	if in.Quota != nil {
		in, out := &in.Quota, &out.Quota
		*out = new(UserQuota)
		(*in).DeepCopyInto(*out)
	}
	if in.Subusers != nil {
		in, out := &in.Subusers, &out.Subusers
		*out = make([]Subuser, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3UserClaimSpec.
func (in *S3UserClaimSpec) DeepCopy() *S3UserClaimSpec {
	if in == nil {
		return nil
	}
	out := new(S3UserClaimSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3UserClaimStatus) DeepCopyInto(out *S3UserClaimStatus) {
	*out = *in
	if in.Quota != nil {
		in, out := &in.Quota, &out.Quota
		*out = new(UserQuota)
		(*in).DeepCopyInto(*out)
	}
	if in.Subusers != nil {
		in, out := &in.Subusers, &out.Subusers
		*out = make([]Subuser, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3UserClaimStatus.
func (in *S3UserClaimStatus) DeepCopy() *S3UserClaimStatus {
	if in == nil {
		return nil
	}
	out := new(S3UserClaimStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3UserList) DeepCopyInto(out *S3UserList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]S3User, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3UserList.
func (in *S3UserList) DeepCopy() *S3UserList {
	if in == nil {
		return nil
	}
	out := new(S3UserList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S3UserList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3UserSpec) DeepCopyInto(out *S3UserSpec) {
	*out = *in
	if in.Quota != nil {
		in, out := &in.Quota, &out.Quota
		*out = new(UserQuota)
		(*in).DeepCopyInto(*out)
	}
	if in.ClaimRef != nil {
		in, out := &in.ClaimRef, &out.ClaimRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3UserSpec.
func (in *S3UserSpec) DeepCopy() *S3UserSpec {
	if in == nil {
		return nil
	}
	out := new(S3UserSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3UserStatus) DeepCopyInto(out *S3UserStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3UserStatus.
func (in *S3UserStatus) DeepCopy() *S3UserStatus {
	if in == nil {
		return nil
	}
	out := new(S3UserStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubuserBinding) DeepCopyInto(out *SubuserBinding) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubuserBinding.
func (in *SubuserBinding) DeepCopy() *SubuserBinding {
	if in == nil {
		return nil
	}
	out := new(SubuserBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserQuota) DeepCopyInto(out *UserQuota) {
	*out = *in
	out.MaxSize = in.MaxSize.DeepCopy()
	out.MaxObjects = in.MaxObjects.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserQuota.
func (in *UserQuota) DeepCopy() *UserQuota {
	if in == nil {
		return nil
	}
	out := new(UserQuota)
	in.DeepCopyInto(out)
	return out
}

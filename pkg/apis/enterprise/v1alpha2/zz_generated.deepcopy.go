// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha2

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonSpec) DeepCopyInto(out *CommonSpec) {
	*out = *in
	in.Affinity.DeepCopyInto(&out.Affinity)
	in.Resources.DeepCopyInto(&out.Resources)
	in.ServiceTemplate.DeepCopyInto(&out.ServiceTemplate)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonSpec.
func (in *CommonSpec) DeepCopy() *CommonSpec {
	if in == nil {
		return nil
	}
	out := new(CommonSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonSplunkSpec) DeepCopyInto(out *CommonSplunkSpec) {
	*out = *in
	in.CommonSpec.DeepCopyInto(&out.CommonSpec)
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.LicenseMasterRef = in.LicenseMasterRef
	out.IndexerRef = in.IndexerRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonSplunkSpec.
func (in *CommonSplunkSpec) DeepCopy() *CommonSplunkSpec {
	if in == nil {
		return nil
	}
	out := new(CommonSplunkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Indexer) DeepCopyInto(out *Indexer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Indexer.
func (in *Indexer) DeepCopy() *Indexer {
	if in == nil {
		return nil
	}
	out := new(Indexer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Indexer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexerList) DeepCopyInto(out *IndexerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Indexer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexerList.
func (in *IndexerList) DeepCopy() *IndexerList {
	if in == nil {
		return nil
	}
	out := new(IndexerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IndexerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexerSpec) DeepCopyInto(out *IndexerSpec) {
	*out = *in
	in.CommonSplunkSpec.DeepCopyInto(&out.CommonSplunkSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexerSpec.
func (in *IndexerSpec) DeepCopy() *IndexerSpec {
	if in == nil {
		return nil
	}
	out := new(IndexerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexerStatus) DeepCopyInto(out *IndexerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexerStatus.
func (in *IndexerStatus) DeepCopy() *IndexerStatus {
	if in == nil {
		return nil
	}
	out := new(IndexerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseMaster) DeepCopyInto(out *LicenseMaster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseMaster.
func (in *LicenseMaster) DeepCopy() *LicenseMaster {
	if in == nil {
		return nil
	}
	out := new(LicenseMaster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LicenseMaster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseMasterList) DeepCopyInto(out *LicenseMasterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LicenseMaster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseMasterList.
func (in *LicenseMasterList) DeepCopy() *LicenseMasterList {
	if in == nil {
		return nil
	}
	out := new(LicenseMasterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LicenseMasterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseMasterSpec) DeepCopyInto(out *LicenseMasterSpec) {
	*out = *in
	in.CommonSplunkSpec.DeepCopyInto(&out.CommonSplunkSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseMasterSpec.
func (in *LicenseMasterSpec) DeepCopy() *LicenseMasterSpec {
	if in == nil {
		return nil
	}
	out := new(LicenseMasterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseMasterStatus) DeepCopyInto(out *LicenseMasterStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseMasterStatus.
func (in *LicenseMasterStatus) DeepCopy() *LicenseMasterStatus {
	if in == nil {
		return nil
	}
	out := new(LicenseMasterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchHead) DeepCopyInto(out *SearchHead) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchHead.
func (in *SearchHead) DeepCopy() *SearchHead {
	if in == nil {
		return nil
	}
	out := new(SearchHead)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SearchHead) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchHeadList) DeepCopyInto(out *SearchHeadList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SearchHead, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchHeadList.
func (in *SearchHeadList) DeepCopy() *SearchHeadList {
	if in == nil {
		return nil
	}
	out := new(SearchHeadList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SearchHeadList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchHeadMemberStatus) DeepCopyInto(out *SearchHeadMemberStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchHeadMemberStatus.
func (in *SearchHeadMemberStatus) DeepCopy() *SearchHeadMemberStatus {
	if in == nil {
		return nil
	}
	out := new(SearchHeadMemberStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchHeadSpec) DeepCopyInto(out *SearchHeadSpec) {
	*out = *in
	in.CommonSplunkSpec.DeepCopyInto(&out.CommonSplunkSpec)
	out.SparkRef = in.SparkRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchHeadSpec.
func (in *SearchHeadSpec) DeepCopy() *SearchHeadSpec {
	if in == nil {
		return nil
	}
	out := new(SearchHeadSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchHeadStatus) DeepCopyInto(out *SearchHeadStatus) {
	*out = *in
	if in.Members != nil {
		in, out := &in.Members, &out.Members
		*out = make([]SearchHeadMemberStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchHeadStatus.
func (in *SearchHeadStatus) DeepCopy() *SearchHeadStatus {
	if in == nil {
		return nil
	}
	out := new(SearchHeadStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Spark) DeepCopyInto(out *Spark) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Spark.
func (in *Spark) DeepCopy() *Spark {
	if in == nil {
		return nil
	}
	out := new(Spark)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Spark) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SparkList) DeepCopyInto(out *SparkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Spark, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SparkList.
func (in *SparkList) DeepCopy() *SparkList {
	if in == nil {
		return nil
	}
	out := new(SparkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SparkList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SparkSpec) DeepCopyInto(out *SparkSpec) {
	*out = *in
	in.CommonSpec.DeepCopyInto(&out.CommonSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SparkSpec.
func (in *SparkSpec) DeepCopy() *SparkSpec {
	if in == nil {
		return nil
	}
	out := new(SparkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SparkStatus) DeepCopyInto(out *SparkStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SparkStatus.
func (in *SparkStatus) DeepCopy() *SparkStatus {
	if in == nil {
		return nil
	}
	out := new(SparkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Standalone) DeepCopyInto(out *Standalone) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Standalone.
func (in *Standalone) DeepCopy() *Standalone {
	if in == nil {
		return nil
	}
	out := new(Standalone)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Standalone) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StandaloneList) DeepCopyInto(out *StandaloneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Standalone, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StandaloneList.
func (in *StandaloneList) DeepCopy() *StandaloneList {
	if in == nil {
		return nil
	}
	out := new(StandaloneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StandaloneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StandaloneSpec) DeepCopyInto(out *StandaloneSpec) {
	*out = *in
	in.CommonSplunkSpec.DeepCopyInto(&out.CommonSplunkSpec)
	out.SparkRef = in.SparkRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StandaloneSpec.
func (in *StandaloneSpec) DeepCopy() *StandaloneSpec {
	if in == nil {
		return nil
	}
	out := new(StandaloneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StandaloneStatus) DeepCopyInto(out *StandaloneStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StandaloneStatus.
func (in *StandaloneStatus) DeepCopy() *StandaloneStatus {
	if in == nil {
		return nil
	}
	out := new(StandaloneStatus)
	in.DeepCopyInto(out)
	return out
}

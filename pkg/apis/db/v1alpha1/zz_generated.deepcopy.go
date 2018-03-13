// +build !ignore_autogenerated

/*
Copyright 2018 Infoblox, Inc.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Database) DeepCopyInto(out *Database) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Database.
func (in *Database) DeepCopy() *Database {
	if in == nil {
		return nil
	}
	out := new(Database)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Database) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseList) DeepCopyInto(out *DatabaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Database, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseList.
func (in *DatabaseList) DeepCopy() *DatabaseList {
	if in == nil {
		return nil
	}
	out := new(DatabaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatabaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseServer) DeepCopyInto(out *DatabaseServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseServer.
func (in *DatabaseServer) DeepCopy() *DatabaseServer {
	if in == nil {
		return nil
	}
	out := new(DatabaseServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatabaseServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseServerList) DeepCopyInto(out *DatabaseServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DatabaseServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseServerList.
func (in *DatabaseServerList) DeepCopy() *DatabaseServerList {
	if in == nil {
		return nil
	}
	out := new(DatabaseServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatabaseServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseServerPlugin) DeepCopyInto(out *DatabaseServerPlugin) {
	*out = *in
	if in.RDS != nil {
		in, out := &in.RDS, &out.RDS
		if *in == nil {
			*out = nil
		} else {
			*out = new(RDSPlugin)
			**out = **in
		}
	}
	if in.MySQL != nil {
		in, out := &in.MySQL, &out.MySQL
		if *in == nil {
			*out = nil
		} else {
			*out = new(MySQLPlugin)
			**out = **in
		}
	}
	if in.Postgres != nil {
		in, out := &in.Postgres, &out.Postgres
		if *in == nil {
			*out = nil
		} else {
			*out = new(PostgresPlugin)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseServerPlugin.
func (in *DatabaseServerPlugin) DeepCopy() *DatabaseServerPlugin {
	if in == nil {
		return nil
	}
	out := new(DatabaseServerPlugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseServerSpec) DeepCopyInto(out *DatabaseServerSpec) {
	*out = *in
	in.DatabaseServerPlugin.DeepCopyInto(&out.DatabaseServerPlugin)
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SuperUserFrom != nil {
		in, out := &in.SuperUserFrom, &out.SuperUserFrom
		if *in == nil {
			*out = nil
		} else {
			*out = new(ValueSource)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.SuperUserPasswordFrom != nil {
		in, out := &in.SuperUserPasswordFrom, &out.SuperUserPasswordFrom
		if *in == nil {
			*out = nil
		} else {
			*out = new(ValueSource)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseServerSpec.
func (in *DatabaseServerSpec) DeepCopy() *DatabaseServerSpec {
	if in == nil {
		return nil
	}
	out := new(DatabaseServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseServerStatus) DeepCopyInto(out *DatabaseServerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseServerStatus.
func (in *DatabaseServerStatus) DeepCopy() *DatabaseServerStatus {
	if in == nil {
		return nil
	}
	out := new(DatabaseServerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseSpec) DeepCopyInto(out *DatabaseSpec) {
	*out = *in
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]DatabaseUser, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DsnFrom != nil {
		in, out := &in.DsnFrom, &out.DsnFrom
		if *in == nil {
			*out = nil
		} else {
			*out = new(ValueSource)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseSpec.
func (in *DatabaseSpec) DeepCopy() *DatabaseSpec {
	if in == nil {
		return nil
	}
	out := new(DatabaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseStatus) DeepCopyInto(out *DatabaseStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseStatus.
func (in *DatabaseStatus) DeepCopy() *DatabaseStatus {
	if in == nil {
		return nil
	}
	out := new(DatabaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseUser) DeepCopyInto(out *DatabaseUser) {
	*out = *in
	if in.PasswordFrom != nil {
		in, out := &in.PasswordFrom, &out.PasswordFrom
		if *in == nil {
			*out = nil
		} else {
			*out = new(ValueSource)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseUser.
func (in *DatabaseUser) DeepCopy() *DatabaseUser {
	if in == nil {
		return nil
	}
	out := new(DatabaseUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLPlugin) DeepCopyInto(out *MySQLPlugin) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLPlugin.
func (in *MySQLPlugin) DeepCopy() *MySQLPlugin {
	if in == nil {
		return nil
	}
	out := new(MySQLPlugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresPlugin) DeepCopyInto(out *PostgresPlugin) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresPlugin.
func (in *PostgresPlugin) DeepCopy() *PostgresPlugin {
	if in == nil {
		return nil
	}
	out := new(PostgresPlugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDSPlugin) DeepCopyInto(out *RDSPlugin) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDSPlugin.
func (in *RDSPlugin) DeepCopy() *RDSPlugin {
	if in == nil {
		return nil
	}
	out := new(RDSPlugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValueSource) DeepCopyInto(out *ValueSource) {
	*out = *in
	if in.ConfigMapKeyRef != nil {
		in, out := &in.ConfigMapKeyRef, &out.ConfigMapKeyRef
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.ConfigMapKeySelector)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.SecretKeyRef != nil {
		in, out := &in.SecretKeyRef, &out.SecretKeyRef
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.SecretKeySelector)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValueSource.
func (in *ValueSource) DeepCopy() *ValueSource {
	if in == nil {
		return nil
	}
	out := new(ValueSource)
	in.DeepCopyInto(out)
	return out
}

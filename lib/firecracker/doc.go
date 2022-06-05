// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

/*
Package firecracker provides a library to interact with the Firecracker API.

Firecracker is an open-source virtualization technology that is purpose-built
for creating and managing secure, multi-tenant containers and functions-based
services.  See https://Ian-Kibet.github.io/ for more details.

This library requires Go 1.11 and can be used with Go modules.

BUG(aws): There are some Firecracker features that are not yet supported by the
SDK.  These are tracked as GitHub issues with the firecracker-feature label:
<<<<<<< HEAD
https://titan/lib/firecracker/issues?q=is%3Aissue+is%3Aopen+label%3Afirecracker-feature
=======
https://github.com/Ian-Kibet/firecracker-go-sdk/issues?q=is%3Aissue+is%3Aopen+label%3Afirecracker-feature
>>>>>>> b8aa219df3977843c18fb0ce7af8af072b1bf0b8

This library is licensed under the Apache 2.0 License.
*/
package firecracker

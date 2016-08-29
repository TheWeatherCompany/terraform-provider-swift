# terraform-provider-swift

## Install

```
$ go get github.com/TheWeatherCompany/terraform-provider-swift
```

Create or edit this file to specify the location of the terraform swift provider binary:

```
# ~/.terraformrc
providers {
    swift = "/path/to/bin/terraform-provider-swift"
}
```

## Documentation

### Provider

The Swift provider is used to interact directly with a Swift-compatible object store.

Use the navigation to the left to read about the available resources.

<div class="alert alert-block alert-info">
<strong>Note:</strong> The Swift provider is brand new.
It is ready to be used but many features are still being added. If there
is a Swift feature missing, please report it in the GitHub repo.
</div>

#### Example Usage

Here is an example that will setup the following:

+ A container.
+ An object within that container.

(create this as myswift.tf and run terraform commands from this directory):

```hcl
provider "swift" {
    username = "" # The user name to use for Swift API operations.
    api_key = ""  # The API key to use for Swift API operations.
    auth_url = "" # The swifth object storage url to use for authentication.
    # Optional. Alternate object storage url to access containers in
    # (defaults to storage url returned by authentication api).
    storage_url = ""
}

# This will create a new container for object in the Swift object store.
resource "swift_container" "test_container_1" {
    name = "test_container_1"
}

# This will create an object under the specified container in
# the Swift object store.
resource "swift_object" "test_object_1" {
    name = "foo.txt" # Object name
    container_name = "${swift_container.test_container_1.name}"
    contents = "${file("foo.txt")}" # Contents of the new object
}
```

You'll need to provide your Swift username, API key, and authentication endpoint,
so that Terraform can create a connection. If you don't want to put
credentials in your configuration file, you can leave them
out:

```
provider "swift" {}
```

...and instead set these environment variables:

- **SWIFT_USERNAME**: Your Swift username
- **SWIFT_API_KEY**: Your Swift API key
- **SWIFT_AUTH_URL**: Your Swift authentication endpoint URL

### Resources

#### swift_container

Creates an object container in a Swift object store.

##### Example Usage

```
resource "swift_container" "test_container" {
    name = "test_containe_name"
    read_access = ["joe", "pete"]
    write_access = ["bob"]
}
```

##### Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the object container.
* `read_access` - (Optional) A list of usernames that will have read access to objects in this container.
* `write_access` - (Optional) A list of usernames that will have write access to objects in this container.

Fields `read_access` and `write_access` are editable.

#### swift_object

Provides an object resource. Allows the creation of an object in a Swift object store.

##### Example Usage

```
# Create a new object in Swift with contents extracted from a local file
resource "swift_object" "test_object" {
    name = "foo.txt" # Object name
    container_name = "${swift_container.test_container_1.name}"
    contents = "${file("foo.txt")}" # Contents of the new object
}
```

```
# Create a new object in Swift with contents specified as a variable.
# NOTE: the path specified will automatically be created
variable "secrets" {
    type = "string"
}

resource "swift_object" "test_object2" {
    name = "path/bar.txt" # Path name/Object name
    container_name = "${swift_container.test_container_1.name}"
    contents = "${var.secrets}"
}
```

##### Argument Reference

The following arguments are supported:

* `name` | *string*
    * Name of the object. This name can also have forward slashes, which will act as a pseudo file path identifying the object location within the container (e.g. _path/to/foo.txt_).
    * **Required**
* `container_name` | *string*
    * Name of the container to put this object in.
    * **Required**
* `contents` | *string*
    * The desired contents of the object.
    * **Optional**

If `contents` is not specified, an empty object will be created.

## Development

### Setup

You should have the correct source in your _$GOPATH_ for both terraform and github.com/ncw/swift.

To get the swift library:

```
go get github.com/ncw/swift
```

To get _terraform_:

```
cd $GOPATH/src
mkdir -p github.com/hashicorp
cd github.com/hashicorp
git clone git@github.com:hashicorp/terraform.git
```

### Build

```
make bin
```

### Test

```
make test
```

### Updating dependencies

We are using [govendor](https://github.com/kardianos/govendor) to manage dependencies just like Terraform. Please see its documentation for additional help.

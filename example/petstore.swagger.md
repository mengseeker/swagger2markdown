# Swagger Petstore (1.0.7)

**Schemes:** https http 

---

## POST /pet
Add a new pet to the store
### Parameters
#### Body parameters
| Name | type | Required | Description |
| ---- |  ---- | -------- | ----------- |
| status | string | false | pet status in the store |
| id | integer(int64) | false |  |
| category | object | false |  |
| category.id | integer(int64) | false |  |
| category.name | string | false |  |
| name | string | true |  |
| photoUrls | array | true |  |
| photoUrls[] | string | true |  |
| tags | array | false |  |
| tags[] | object | false |  |
| tags[].id | integer(int64) | false |  |
| tags[].name | string | false |  |

### Responses

#### 405
Invalid input

## PUT /pet
Update an existing pet
### Parameters
#### Body parameters
| Name | type | Required | Description |
| ---- |  ---- | -------- | ----------- |
| id | integer(int64) | false |  |
| category | object | false |  |
| category.name | string | false |  |
| category.id | integer(int64) | false |  |
| name | string | true |  |
| photoUrls | array | true |  |
| photoUrls[] | string | true |  |
| tags | array | false |  |
| tags[] | object | false |  |
| tags[].id | integer(int64) | false |  |
| tags[].name | string | false |  |
| status | string | false | pet status in the store |

### Responses

#### 400
Invalid ID supplied

#### 404
Pet not found

#### 405
Validation exception

## GET /pet/findByStatus
Finds Pets by status
### Parameters
#### Query parameters
| Name | type | Required | Description |
| ---- |  ---- | -------- | ----------- |
| status | []string | true | Status values that need to be considered for filter |

### Responses

#### 200
successful operation
| Name | type | Required | Description |
| ---- | ---- | -------- | ----------- |
| [] | object | true |  |
| [].status | string | false | pet status in the store |
| [].id | integer(int64) | false |  |
| [].category | object | false |  |
| [].category.id | integer(int64) | false |  |
| [].category.name | string | false |  |
| [].name | string | true |  |
| [].photoUrls | array | true |  |
| [].photoUrls[] | string | true |  |
| [].tags | array | false |  |
| [].tags[] | object | false |  |
| [].tags[].id | integer(int64) | false |  |
| [].tags[].name | string | false |  |

#### 400
Invalid status value

## GET /pet/findByTags
Finds Pets by tags
### Parameters
#### Query parameters
| Name | type | Required | Description |
| ---- |  ---- | -------- | ----------- |
| tags | []string | true | Tags to filter by |

### Responses

#### 200
successful operation
| Name | type | Required | Description |
| ---- | ---- | -------- | ----------- |
| [] | object | true |  |
| [].tags | array | false |  |
| [].tags[] | object | false |  |
| [].tags[].id | integer(int64) | false |  |
| [].tags[].name | string | false |  |
| [].status | string | false | pet status in the store |
| [].id | integer(int64) | false |  |
| [].category | object | false |  |
| [].category.id | integer(int64) | false |  |
| [].category.name | string | false |  |
| [].name | string | true |  |
| [].photoUrls | array | true |  |
| [].photoUrls[] | string | true |  |

#### 400
Invalid tag value

## DELETE /pet/{petId}
Deletes a pet
### Parameters
#### Header parameters
| Name | type | Required | Description |
| ---- |  ---- | -------- | ----------- |
| api_key | string | false |  |
#### Path parameters
| Name | type | Required | Description |
| ---- |  ---- | -------- | ----------- |
| petId | integer | true | Pet id to delete |

### Responses

#### 400
Invalid ID supplied

#### 404
Pet not found

## GET /pet/{petId}
Find pet by ID
### Parameters
#### Path parameters
| Name | type | Required | Description |
| ---- |  ---- | -------- | ----------- |
| petId | integer | true | ID of pet to return |

### Responses

#### 200
successful operation
| Name | type | Required | Description |
| ---- | ---- | -------- | ----------- |
| id | integer(int64) | false |  |
| category | object | false |  |
| category.id | integer(int64) | false |  |
| category.name | string | false |  |
| name | string | true |  |
| photoUrls | array | true |  |
| photoUrls[] | string | true |  |
| tags | array | false |  |
| tags[] | object | false |  |
| tags[].id | integer(int64) | false |  |
| tags[].name | string | false |  |
| status | string | false | pet status in the store |

#### 400
Invalid ID supplied

#### 404
Pet not found

## POST /pet/{petId}
Updates a pet in the store with form data
### Parameters
#### Path parameters
| Name | type | Required | Description |
| ---- |  ---- | -------- | ----------- |
| petId | integer | true | ID of pet that needs to be updated |

### Responses

#### 405
Invalid input

## POST /pet/{petId}/uploadImage
uploads an image
### Parameters
#### Path parameters
| Name | type | Required | Description |
| ---- |  ---- | -------- | ----------- |
| petId | integer | true | ID of pet to update |

### Responses

#### 200
successful operation
| Name | type | Required | Description |
| ---- | ---- | -------- | ----------- |
| code | integer(int32) | false |  |
| type | string | false |  |
| message | string | false |  |

## GET /store/inventory
Returns pet inventories by status

### Responses

#### 200
successful operation

## POST /store/order
Place an order for a pet
### Parameters
#### Body parameters
| Name | type | Required | Description |
| ---- |  ---- | -------- | ----------- |
| id | integer(int64) | false |  |
| petId | integer(int64) | false |  |
| quantity | integer(int32) | false |  |
| shipDate | string(date-time) | false |  |
| status | string | false | Order Status |
| complete | boolean | false |  |

### Responses

#### 200
successful operation
| Name | type | Required | Description |
| ---- | ---- | -------- | ----------- |
| complete | boolean | false |  |
| id | integer(int64) | false |  |
| petId | integer(int64) | false |  |
| quantity | integer(int32) | false |  |
| shipDate | string(date-time) | false |  |
| status | string | false | Order Status |

#### 400
Invalid Order

## DELETE /store/order/{orderId}
Delete purchase order by ID
### Parameters
#### Path parameters
| Name | type | Required | Description |
| ---- |  ---- | -------- | ----------- |
| orderId | integer | true | ID of the order that needs to be deleted |

### Responses

#### 400
Invalid ID supplied

#### 404
Order not found

## GET /store/order/{orderId}
Find purchase order by ID
### Parameters
#### Path parameters
| Name | type | Required | Description |
| ---- |  ---- | -------- | ----------- |
| orderId | integer | true | ID of pet that needs to be fetched |

### Responses

#### 200
successful operation
| Name | type | Required | Description |
| ---- | ---- | -------- | ----------- |
| status | string | false | Order Status |
| complete | boolean | false |  |
| id | integer(int64) | false |  |
| petId | integer(int64) | false |  |
| quantity | integer(int32) | false |  |
| shipDate | string(date-time) | false |  |

#### 400
Invalid ID supplied

#### 404
Order not found

## POST /user
Create user
### Parameters
#### Body parameters
| Name | type | Required | Description |
| ---- |  ---- | -------- | ----------- |
| password | string | false |  |
| phone | string | false |  |
| userStatus | integer(int32) | false | User Status |
| id | integer(int64) | false |  |
| username | string | false |  |
| firstName | string | false |  |
| lastName | string | false |  |
| email | string | false |  |

### Responses

#### default
successful operation

## POST /user/createWithArray
Creates list of users with given input array
### Parameters
#### Body parameters
| Name | type | Required | Description |
| ---- |  ---- | -------- | ----------- |
| [] | object | true |  |
| [].id | integer(int64) | false |  |
| [].username | string | false |  |
| [].firstName | string | false |  |
| [].lastName | string | false |  |
| [].email | string | false |  |
| [].password | string | false |  |
| [].phone | string | false |  |
| [].userStatus | integer(int32) | false | User Status |

### Responses

#### default
successful operation

## POST /user/createWithList
Creates list of users with given input array
### Parameters
#### Body parameters
| Name | type | Required | Description |
| ---- |  ---- | -------- | ----------- |
| [] | object | true |  |
| [].id | integer(int64) | false |  |
| [].username | string | false |  |
| [].firstName | string | false |  |
| [].lastName | string | false |  |
| [].email | string | false |  |
| [].password | string | false |  |
| [].phone | string | false |  |
| [].userStatus | integer(int32) | false | User Status |

### Responses

#### default
successful operation

## GET /user/login
Logs user into the system
### Parameters
#### Query parameters
| Name | type | Required | Description |
| ---- |  ---- | -------- | ----------- |
| username | string | true | The user name for login |
| password | string | true | The password for login in clear text |

### Responses

#### 200
successful operation

#### 400
Invalid username/password supplied

## GET /user/logout
Logs out current logged in user session

### Responses

#### default
successful operation

## DELETE /user/{username}
Delete user
### Parameters
#### Path parameters
| Name | type | Required | Description |
| ---- |  ---- | -------- | ----------- |
| username | string | true | The name that needs to be deleted |

### Responses

#### 400
Invalid username supplied

#### 404
User not found

## GET /user/{username}
Get user by user name
### Parameters
#### Path parameters
| Name | type | Required | Description |
| ---- |  ---- | -------- | ----------- |
| username | string | true | The name that needs to be fetched. Use user1 for testing.  |

### Responses

#### 200
successful operation
| Name | type | Required | Description |
| ---- | ---- | -------- | ----------- |
| userStatus | integer(int32) | false | User Status |
| id | integer(int64) | false |  |
| username | string | false |  |
| firstName | string | false |  |
| lastName | string | false |  |
| email | string | false |  |
| password | string | false |  |
| phone | string | false |  |

#### 400
Invalid username supplied

#### 404
User not found

## PUT /user/{username}
Updated user
### Parameters
#### Path parameters
| Name | type | Required | Description |
| ---- |  ---- | -------- | ----------- |
| username | string | true | name that need to be updated |
#### Body parameters
| Name | type | Required | Description |
| ---- |  ---- | -------- | ----------- |
| email | string | false |  |
| password | string | false |  |
| phone | string | false |  |
| userStatus | integer(int32) | false | User Status |
| id | integer(int64) | false |  |
| username | string | false |  |
| firstName | string | false |  |
| lastName | string | false |  |

### Responses

#### 400
Invalid user supplied

#### 404
User not found

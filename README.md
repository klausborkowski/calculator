# Read.me contents

1. [Overview](#1-overview)
2. [Build Project Information](#2-build-project-information)
3. [Check Test Coverage Information](#3-check-test-coverage-information)
4. [UI](#4-ui)
5. [API](#5-api)
6. [Deployed Service](#6-deployed-service)

---

# Packager Service

## 1. Overview  
Our customers can order any number of these items through our website, but they will always only be given complete packs.

### Order Fulfillment Rules:
- Only whole packs can be sent. Packs cannot be broken open.
- Within the constraints of Rule 1 above, send out the least amount of items to fulfill the order.
- Within the constraints of Rules 1 & 2 above, send out as few packs as possible to fulfill each order.

---

## 2. Build Project Information  
To build and run the project, execute:  
```sh
make docker-up
```
Then open in the browser:  
```
http://localhost:8080/app
```

---

## 3. Check Test Coverage Information  
To check test coverage, run:  
```sh
make test-report
```
This will generate and open a UI report in the browser.

---

## 4. UI  
- **Swagger Button**: Opens Swagger API documentation.  
- **Add Package Sizes**: Use the **X** button to remove a package size and the **✔️** button to submit.  
- **Order Input**: Enter the desired order size in the input field.  
- **Calculate Button**: Click to compute the optimized package distribution.  

---

## 5. API  
The API can be accessed via:  
```
http://localhost:8080/swagger/index.html
```
Alternatively, use the **Swagger Button** in the UI to open the API docs.

## 6. Deployed service
THere is packager deployed publicly here:[Packager Service](https://packager-0e6j.onrender.com/app)
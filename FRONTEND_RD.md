# Frontend Requirements Document

This document outlines the requirements and API specifications for developing the frontend client (React + TypeScript) for the Address Book Server V3.

## 1. Tech Stack Requirements
* **Framework:** React.js
* **Language:** TypeScript
* **State Management:** Redux Toolkit
* **Routing:** React Router
* **Form Handling & Validation:** React Hook Form
* **Styling:** CSS Modules, SCSS
* **HTTP Client:** Axios (Use axios version that is not vulnerable, as on 31st March 2026 vulnerability is found in axios's latest release.)

## 2. API Base Setup
* **Base URL:** `http://localhost:8000/api/v3`
* **Authentication:** JWT (JSON Web Token)
  * Tokens should be stored securely (e.g., in `localStorage` or secure cookies).
  * All protected routes must include the token in the `Authorization` header:
    `Authorization: Bearer <token>`

---

## 3. Data Models

### Address Object
```typescript
interface Address {
  id: string;
  user_id: string;
  first_name: string;
  last_name: string;
  email: string;
  phone?: string;
  address_line1: string;
  address_line2?: string;
  city?: string;
  state?: string;
  country?: string;
  pincode?: string;
}
```

---

## 4. Required Validations (Mapped from Backend)

To ensure smooth API communication, the frontend MUST implement the following validations before submitting requests:

* **Email:** Must be a valid format: `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
* **Password:** Minimum 8 characters. Must contain at least one uppercase letter, one lowercase letter, one digit, and one special character.
* **Phone Number:** Exactly 10 digits, must start with 6, 7, 8, or 9 (Regex: `^[6-9]\d{9}$`).
* **Pincode:** Exactly 6 digits, first digit cannot be 0 (Regex: `^[1-9][0-9]{5}$`).

---

## 5. API Endpoints & Features

### 5.1 Authentication Features

#### **A. Register User**
* **Endpoint:** `POST /auth/register`
* **Payload:**
  ```typescript
  {
    email: string;      // Required, valid email
    password: string;   // Required, strong password
  }
  ```
* **Response:** `{ id: string, email: string }`

#### **B. Login User**
* **Endpoint:** `POST /auth/login`
* **Payload:**
  ```typescript
  {
    email: string;      // Required, valid email
    password: string;   // Required
  }
  ```
* **Response:** `{ token: string }`
* **Action:** Save `token` to manage the user's session.

---

### 5.2 Address Book Features (Protected Routes)

#### **A. Create Address**
* **Endpoint:** `POST /addresses`
* **Payload:**
  ```typescript
  {
    first_name: string;      // Required
    last_name?: string;
    email: string;           // Required, valid email
    phone?: string;          // Valid phone number
    address_line1: string;   // Required
    address_line2?: string;
    city?: string;
    state?: string;
    country?: string;
    pincode?: string;        // Valid pincode
  }
  ```
* **Response:** Returns the created `Address` object.

#### **B. Get All Addresses**
* **Endpoint:** `GET /addresses`
* **Response:**
  ```typescript
  {
    addresses: Address[]
  }
  ```

#### **C. Get Address by ID**
* **Endpoint:** `GET /addresses/:id`
* **Response:** Returns the specific `Address` object.

#### **D. Update Address**
* **Endpoint:** `PUT /addresses/:id`
* **Payload:** All properties from the Create Address payload, but entirely optional. Send only the fields passing edits.
* **Response:** Returns the updated `Address` object.

#### **E. Delete Address**
* **Endpoint:** `DELETE /addresses/:id`
* **Response:** `{ message: string }`

---

### 5.3 Advanced Features

#### **A. Filter/Paginate Addresses**
* **Endpoint:** `GET /addresses/filter`
* **Query Parameters:**
  * `page` (number) - For pagination (Default: `1`)
  * `limit` (number) - Items per page (Default: `10`)
  * `search` (string) - Generic search term
  * `city` (string)
  * `state` (string)
  * `country` (string)
  * `pincode` (string)
* **Response:**
  ```typescript
  {
    data: {
      addresses: Address[]
    },
    total: number // Total matching records for calculating total pages
  }
  ```

#### **B. Export Custom Addresses**
* **Endpoint:** `POST /addresses/export`
* **Description:** Asynchronously prepares a CSV file based on the requested fields and emails a download link to the user.
* **Payload:**
  ```typescript
  {
    fields: string[]; // Required, at least 1 item (e.g., ["first_name", "email", "phone"])
    email: string;    // Required, destination email to receive the download link
  }
  ```
* **Response:** `{ message: "Export started" }` (or similar success message).

---

## 6. Suggested UI Architecture / Views

1. **Authentication:**
   * `/login` - Login Page
   * `/register` - Registration Page
2. **Dashboard / Home:**
   * `/` - Displays a paginated data table or list of addresses (utilizing `GET /addresses/filter`).
   * **Search/Filter Bar:** Incorporate fields to filter by `city`, `state`, `country`, `pincode`, and generic `search`.
3. **Address Management:**
   * `/address/new` - Form to create a new address.
   * `/address/edit/:id` - Form pre-filled with data to edit an existing address.
   * `View Modal/Page` - Dedicated view to see all details of an address using `GET /addresses/:id`.
4. **Export Flow:**
   * A "Export" button on the dashboard allowing users to pick fields (via checkboxes) and specifying their email to receive the report.

## 7. Essential Implementation Topics

### Essential TypeScript Topics
* Type annotations & custom types
* Interfaces & classes
* Enums
* Union and Intersection
* Keyof and typeof operators
* Mapped types

### Essential JS Topics
* **Basics:** declaration and data types (primitive, non-primitive, json, null, undefined)
* **Objects and prototypes:** (key value, entry, comparison, copy)
* **Operators:** (arithmetic, comparison, ternary, logical, unary, string)
* **Array:** (map, filter, reduce)
* **Functions and scopes**
* **Conditional Statements & Expressions:** (for, if-else, switch-case, IIFE)
* **Advance:**
  * Promise, async/await
  * regex
  * ES6+ (let-const, de-structuring, template, arrow functions, default param, optional chaining)
* **Patterns:**
  * callback
  * Closure
  * HOF
  * Recursion
  * Currying

### Functional Programming Essentials in HTML & CSS
#### HTML
* HTML Structure/Layout
* Tags (Basic, Form, Semantic)
* Image, Links, ref, (JS, CSS, font) files
* meta tags

#### CSS
* Basics (properties, types)
* Size and Units
* Box-Model
* Position & Layout
* Other Imp Properties (bg, visibility)
* Selectors
* Responsiveness
* Transform, Gradient & Animation

#### DOM
* Access Element (id, tag, class, node)
* HTML Manipulation
* Attributes (get/set)
* Element Creation
* Event (listener, bubbling, propagation, prevent-default)

### Essentials in React & Redux
#### React
* JSX/TSX
* Components
* Functional Components
* State and props
* Must know Hooks (useRef, useState, useEffect, useReducer, useCallback, useMemo)
* life cycle
* Events
* Context
* Styling & CSS Module
* **Patterns:**
  * Composition ([Composition vs Inheritance](https://legacy.reactjs.org/docs/composition-vs-inheritance.html))
  * HOC

#### Redux
* Redux Concept & Architecture
* Store, Action & Reducer
* Middlewares
* Thunk
* Redux Toolkit
* Slice
* AsyncActions
* DevTool
* **Patterns:**
  * Immutability
  * Pure Functions
  * Adaptor

#### Styled Component
* Creating Styled Component
* Extending Styles
* Props in Styled Component
* Defining common css
* Theme
* Ref

#### Axios
* Making an API Call
* Axios config
* Handling errors
* Interceptors

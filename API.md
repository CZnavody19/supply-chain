# Supply Chain Network — API Reference

> **Base URL:** `http://localhost:8080`
>
> All endpoints return JSON. All IDs are passed as **query parameters** (never in the URL path).
> The backend runs on port **8080** by default.

---

## Table of Contents

1. [Products](#1-products)
2. [Products — BOM (Bill of Materials)](#2-products--bom-bill-of-materials)
3. [Products — Alternative Suppliers](#3-products--alternative-suppliers)
4. [Companies](#4-companies)
5. [Companies — Risk Assessment](#5-companies--risk-assessment)
6. [Components](#6-components)
7. [Orders](#7-orders)
8. [Orders — Supply Path & Cost](#8-orders--supply-path--cost)
9. [Locations](#9-locations)
10. [Locations — Inventory](#10-locations--inventory)
11. [Routes](#11-routes)
12. [Analytics](#12-analytics)
13. [Predictions](#13-predictions)
14. [Seed Data IDs Reference](#seed-data-ids-reference)

---

## 1. Products

### `GET /api/products`

List all products, or get a single product by ID.

| Query Param | Type   | Required | Description                              |
| ----------- | ------ | -------- | ---------------------------------------- |
| `id`        | string | no       | If provided, returns only that product.  |

**Response (list):** `Product[]`
**Response (single):** `Product`

```jsonc
// Product
{
  "id": "prod-001",
  "name": "Laptop Model X",
  "sku": "LAP-X-001",
  "price": 999.99,
  "weight": 2.1,
  "leadTime": 14,       // days
  "status": "active"    // "active" | "discontinued" | ...
}
```

---

### `POST /api/products`

Create a new product.

**Request Body:** `Product` (without `id` — generated server-side)

```json
{
  "name": "Tablet Z",
  "sku": "TAB-Z-001",
  "price": 499.99,
  "weight": 0.5,
  "leadTime": 10,
  "status": "active"
}
```

**Response:** `201 Created` — the created `Product` with generated `id`.

---

### `PUT /api/products`

Update an existing product.

| Query Param | Type   | Required | Description              |
| ----------- | ------ | -------- | ------------------------ |
| `id`        | string | **yes**  | The product ID to update |

**Request Body:** `Product` fields to update (the `id` in the body is ignored; the query param is used).

**Response:** `200 OK` — the updated `Product`.

---

### `DELETE /api/products`

Delete a product and its relationships.

| Query Param | Type   | Required | Description              |
| ----------- | ------ | -------- | ------------------------ |
| `id`        | string | **yes**  | The product ID to delete |

**Response:** `204 No Content`

---

## 2. Products — BOM (Bill of Materials)

### `GET /api/products/bom`

Get the bill of materials for a product (component list with quantities).

| Query Param | Type   | Required | Description  |
| ----------- | ------ | -------- | ------------ |
| `id`        | string | **yes**  | Product ID   |

**Response:** `BOMEntry[]`

```jsonc
// BOMEntry
{
  "component": { /* Component object */ },
  "quantity": 2,
  "position": 1
}
```

---

### `GET /api/products/bom/detailed`

Same as above but includes supplier info for each component.

| Query Param | Type   | Required | Description  |
| ----------- | ------ | -------- | ------------ |
| `id`        | string | **yes**  | Product ID   |

**Response:** `BOMDetailedEntry[]`

```jsonc
// BOMDetailedEntry
{
  "component": { /* Component */ },
  "quantity": 2,
  "position": 1,
  "suppliers": [
    {
      "company": { /* Company */ },
      "price": 120.0,
      "leadTime": 10,
      "minOrder": 100
    }
  ]
}
```

---

### `POST /api/products/bom`

Add a component to a product's BOM.

| Query Param | Type   | Required | Description  |
| ----------- | ------ | -------- | ------------ |
| `id`        | string | **yes**  | Product ID   |

**Request Body:**

```json
{
  "componentId": "comp-001",
  "quantity": 2,
  "position": 1
}
```

**Response:** `201 Created`

---

### `PUT /api/products/bom`

Update the quantity of a component in a product's BOM.

| Query Param   | Type   | Required | Description   |
| ------------- | ------ | -------- | ------------- |
| `id`          | string | **yes**  | Product ID    |
| `componentId` | string | **yes**  | Component ID  |

**Request Body:**

```json
{
  "quantity": 5
}
```

**Response:** `200 OK`

---

## 3. Products — Alternative Suppliers

### `GET /api/products/alternative-suppliers`

Find alternative suppliers for all components of a product.

| Query Param | Type   | Required | Description  |
| ----------- | ------ | -------- | ------------ |
| `id`        | string | **yes**  | Product ID   |

**Response:** `AlternativeSupplier[]`

```jsonc
{
  "company": { /* Company */ },
  "price": 135.0,
  "reliability": 0.85,
  "leadTime": 14
}
```

---

## 4. Companies

### `GET /api/companies`

List all companies, or get a single company by ID. Companies have types like `supplier`, `manufacturer`, `distributor`, `retailer`, `customer`.

| Query Param | Type   | Required | Description                              |
| ----------- | ------ | -------- | ---------------------------------------- |
| `id`        | string | no       | If provided, returns only that company.  |

**Response (list):** `Company[]`
**Response (single):** `Company`

```jsonc
// Company
{
  "id": "c-supplier-01",
  "name": "ChipCo Taiwan",
  "type": "supplier",
  "country": "Taiwan",
  "coordinates": { "lat": 25.033, "lng": 121.565 },
  "reliability": 0.95    // 0.0 – 1.0
}
```

---

### `POST /api/companies`

Create a new company.

**Request Body:** `Company` (without `id`)

```json
{
  "name": "New Supplier Ltd",
  "type": "supplier",
  "country": "Japan",
  "coordinates": { "lat": 35.689, "lng": 139.692 },
  "reliability": 0.90
}
```

**Response:** `201 Created` — the created `Company`.

---

### `PUT /api/companies`

Update a company.

| Query Param | Type   | Required | Description              |
| ----------- | ------ | -------- | ------------------------ |
| `id`        | string | **yes**  | The company ID to update |

**Request Body:** `Company` fields to update.

**Response:** `200 OK` — the updated `Company`.

---

### `DELETE /api/companies`

Delete a company and its relationships.

| Query Param | Type   | Required | Description              |
| ----------- | ------ | -------- | ------------------------ |
| `id`        | string | **yes**  | The company ID to delete |

**Response:** `204 No Content`

---

## 5. Companies — Risk Assessment

### `GET /api/companies/risk-assessment`

Get a risk assessment for a specific company (supplier).

| Query Param | Type   | Required | Description  |
| ----------- | ------ | -------- | ------------ |
| `id`        | string | **yes**  | Company ID   |

**Response:** `RiskAssessment`

```jsonc
{
  "supplierId": "c-supplier-01",
  "company": "ChipCo Taiwan",
  "riskScore": 25.0,           // 0–100, lower is better
  "factors": {
    "reliabilityScore": 0.95,
    "onTimeDeliveryRate": 0.8,
    "qualityIssues": 0.0,
    "geopoliticalRisk": 0.5,
    "financialStability": 0.7
  },
  "criticalFor": [
    {
      "product": "Laptop Model X",
      "impact": "high",
      "alternatives": 1
    }
  ],
  "recommendations": ["..."]
}
```

---

## 6. Components

### `GET /api/components`

List all components, or get one by ID.

| Query Param | Type   | Required | Description                                |
| ----------- | ------ | -------- | ------------------------------------------ |
| `id`        | string | no       | If provided, returns only that component.  |

**Response:** `Component` or `Component[]`

```jsonc
// Component
{
  "id": "comp-001",
  "name": "CPU Chip",
  "price": 150.0,
  "quantity": 0,
  "criticality": "high"   // "high" | "medium" | "low"
}
```

---

### `POST /api/components`

Create a component.

**Request Body:** `Component` (without `id`)

**Response:** `201 Created` — the created `Component`.

---

### `PUT /api/components`

Update a component.

| Query Param | Type   | Required | Description                |
| ----------- | ------ | -------- | -------------------------- |
| `id`        | string | **yes**  | The component ID to update |

**Request Body:** `Component` fields to update.

**Response:** `200 OK` — the updated `Component`.

---

### `DELETE /api/components`

Delete a component.

| Query Param | Type   | Required | Description                |
| ----------- | ------ | -------- | -------------------------- |
| `id`        | string | **yes**  | The component ID to delete |

**Response:** `204 No Content`

---

## 7. Orders

### `GET /api/orders`

List all orders, or get one by ID.

| Query Param | Type   | Required | Description                            |
| ----------- | ------ | -------- | -------------------------------------- |
| `id`        | string | no       | If provided, returns only that order.  |

**Response:** `Order` or `Order[]`

```jsonc
// Order
{
  "id": "order-001",
  "orderDate": "2024-02-01",
  "dueDate": "2024-02-28",
  "quantity": 100,
  "status": "in_transit",   // "pending" | "in_transit" | "delivered" | "delayed"
  "cost": 45000.0
}
```

---

### `POST /api/orders`

Create a new order with relationships to a product, customer, and supplier.

**Request Body:**

```json
{
  "orderDate": "2024-03-01",
  "dueDate": "2024-03-15",
  "quantity": 50,
  "status": "pending",
  "cost": 25000.0,
  "productId": "prod-001",
  "productQuantity": 50,
  "unitPrice": 450.0,
  "customerId": "c-customer-01",
  "supplierId": "c-manufacturer-01"
}
```

| Body Field        | Type   | Description                                       |
| ----------------- | ------ | ------------------------------------------------- |
| `orderDate`       | string | Order creation date (YYYY-MM-DD)                  |
| `dueDate`         | string | Expected delivery date (YYYY-MM-DD)               |
| `quantity`        | int    | Total quantity                                     |
| `status`          | string | Initial status                                     |
| `cost`            | float  | Total order cost                                   |
| `productId`       | string | ID of the product being ordered                    |
| `productQuantity` | int    | Quantity of the product in this order              |
| `unitPrice`       | float  | Price per unit                                     |
| `customerId`      | string | Company ID of the customer placing the order       |
| `supplierId`      | string | Company ID of the supplier fulfilling the order    |

**Response:** `201 Created` — the created `Order`.

---

### `PUT /api/orders/status`

Update only the status of an order.

| Query Param | Type   | Required | Description  |
| ----------- | ------ | -------- | ------------ |
| `id`        | string | **yes**  | Order ID     |

**Request Body:**

```json
{
  "status": "delivered"
}
```

**Response:** `200 OK`

---

## 8. Orders — Supply Path & Cost

### `GET /api/orders/supply-path`

Trace the full supply path (manufacturing → transport → delivery) for an order.

| Query Param | Type   | Required | Description  |
| ----------- | ------ | -------- | ------------ |
| `orderId`   | string | **yes**  | Order ID     |

**Response:** `SupplyPathResponse`

```jsonc
{
  "orderId": "order-001",
  "product": "Laptop Model X",
  "quantity": 100,
  "totalCost": 45000.0,
  "path": [
    {
      "stage": 1,
      "name": "Manufacturing",
      "company": { "id": "...", "name": "...", "reliability": 0.93 },
      "location": { "id": "...", "name": "..." },
      "dueDate": "2024-02-28",
      "status": "in_transit"
    },
    {
      "stage": 2,
      "name": "Transport",
      "from": "loc-cn-01",
      "to": "loc-de-01",
      "route": { "distance": 10000, "time": "14h", "cost": 8000 },
      "dueDate": "2024-02-28",
      "status": "in_transit"
    },
    {
      "stage": 3,
      "name": "Delivery",
      "company": { "id": "...", "name": "...", "reliability": 1.0 },
      "dueDate": "2024-02-28",
      "status": "in_transit"
    }
  ],
  "totalDuration": "~14 days",
  "riskFactors": ["Supplier reliability: 0.93", "..."]
}
```

---

### `GET /api/orders/cost-breakdown`

Break down the costs of an order into material, manufacturing, and logistics.

| Query Param | Type   | Required | Description  |
| ----------- | ------ | -------- | ------------ |
| `orderId`   | string | **yes**  | Order ID     |

**Response:** `CostBreakdown`

```jsonc
{
  "orderId": "order-001",
  "materialCost": 24000.0,
  "manufacturingCost": 45000.0,
  "logisticsCost": 3000.0,
  "totalCost": 72000.0
}
```

---

## 9. Locations

### `GET /api/locations`

List all locations, or get one by ID.

| Query Param | Type   | Required | Description                               |
| ----------- | ------ | -------- | ----------------------------------------- |
| `id`        | string | no       | If provided, returns only that location.  |

**Response:** `Location` or `Location[]`

```jsonc
// Location
{
  "id": "loc-tw-01",
  "name": "Taiwan Fab",
  "type": "warehouse",    // "warehouse" | "distribution_center" | "port"
  "coordinates": { "lat": 25.033, "lng": 121.565 },
  "capacity": 5000
}
```

---

### `POST /api/locations`

Create a new location.

**Request Body:** `Location` (without `id`)

**Response:** `201 Created` — the created `Location`.

---

## 10. Locations — Inventory

### `GET /api/locations/inventory-status`

Get the inventory status of a location (which products are stored there and how much).

| Query Param | Type   | Required | Description  |
| ----------- | ------ | -------- | ------------ |
| `id`        | string | **yes**  | Location ID  |

**Response:** `InventoryStatus`

```jsonc
{
  "location": { /* Location */ },
  "products": [
    {
      "product": { /* Product */ },
      "quantity": 500,
      "daysOfSupply": 30
    }
  ]
}
```

---

## 11. Routes

### `GET /api/routes/optimal`

Find the shortest (optimal) route between two locations using the graph of connected locations.

| Query Param | Type   | Required | Description              |
| ----------- | ------ | -------- | ------------------------ |
| `from`      | string | **yes**  | Source location ID       |
| `to`        | string | **yes**  | Destination location ID  |

**Response:** `OptimalRouteResult`

```jsonc
{
  "segments": [
    {
      "from": "Taiwan Fab",
      "to": "Singapore Port",
      "distance": 2500,
      "time": 120,
      "cost": 3000
    },
    {
      "from": "Singapore Port",
      "to": "Frankfurt Hub",
      "distance": 10000,
      "time": 14,
      "cost": 8000
    }
  ],
  "totalDistance": 12500,
  "totalTime": 134,
  "totalCost": 11000,
  "totalReliability": 0.0    // product of segment reliabilities if available
}
```

---

## 12. Analytics

### `GET /api/analytics/supply-chain-health`

Overall supply chain health overview: critical components, bottleneck locations, high-risk suppliers.

| Query Param | none | | |
| ----------- | ---- | - | - |

**Response:** `SupplyChainHealth`

```jsonc
{
  "criticalComponents": [
    {
      "componentId": "comp-001",
      "componentName": "CPU Chip",
      "criticality": "high",
      "supplierCount": 2
    }
  ],
  "bottlenecks": [
    {
      "locationId": "loc-de-01",
      "locationName": "Frankfurt Hub",
      "utilization": 0.10
    }
  ],
  "highRiskSuppliers": [
    {
      "companyId": "c-supplier-03",
      "companyName": "ChipCo Vietnam",
      "reliability": 0.85
    }
  ],
  "recommendations": ["Consider adding backup suppliers for high-criticality components with fewer than 2 suppliers"]
}
```

---

### `GET /api/analytics/impact-analysis`

Analyse the downstream impact if a specific supplier were to fail.

| Query Param | Type   | Required | Description                      |
| ----------- | ------ | -------- | -------------------------------- |
| `supplier`  | string | **yes**  | Company ID of the supplier       |

**Response:** `ImpactAnalysis`

```jsonc
{
  "supplierId": "c-supplier-01",
  "supplierName": "ChipCo Taiwan",
  "impact": {
    "affectedProducts": [
      {
        "productId": "prod-001",
        "productName": "Laptop Model X",
        "affectedOrders": 2,
        "delayDays": 10
      }
    ],
    "estimatedCost": 90000.0,
    "affectedRevenue": 199998.0,
    "mitigation": ["Consider alternative supplier: ChipCo Vietnam"]
  }
}
```

---

## 13. Predictions

### `GET /api/analytics/forecast-delays`

Simple heuristic-based delay forecast per product, based on historical order data.

| Query Param | Type | Required | Description                               |
| ----------- | ---- | -------- | ----------------------------------------- |
| `months`    | int  | no       | Forecast horizon in months (default: `3`) |

**Response:** `ForecastDelay[]`

```jsonc
[
  {
    "productId": "prod-001",
    "productName": "Laptop Model X",
    "avgDelayDays": 5.0,
    "probability": 0.4,    // chance of delay in given horizon
    "riskLevel": "medium"  // "low" | "medium" | "high"
  }
]
```

---

### `GET /api/analytics/stock-levels`

Project stock levels for a product over a number of months using linear depletion.

| Query Param | Type   | Required | Description                                       |
| ----------- | ------ | -------- | ------------------------------------------------- |
| `product`   | string | **yes**  | Product ID                                        |
| `months`    | int    | no       | Projection horizon in months (default: `6`)       |

**Response:** `StockLevelForecast`

```jsonc
{
  "productId": "prod-001",
  "productName": "Laptop Model X",
  "currentStock": 700,
  "projections": [
    { "month": "2024-03", "projectedStock": 583, "status": "ok" },
    { "month": "2024-04", "projectedStock": 466, "status": "ok" },
    { "month": "2024-05", "projectedStock": 349, "status": "warning" },
    { "month": "2024-06", "projectedStock": 232, "status": "warning" },
    { "month": "2024-07", "projectedStock": 115, "status": "critical" },
    { "month": "2024-08", "projectedStock": 0,   "status": "critical" }
  ]
}
```

---

## Error Responses

All endpoints return plain-text error messages with standard HTTP status codes:

| Code  | Meaning                                            |
| ----- | -------------------------------------------------- |
| `400` | Bad Request — missing required query param or body |
| `404` | Not Found — entity with given ID doesn't exist     |
| `500` | Internal Server Error                              |

---

## Seed Data IDs Reference

Use these IDs when testing against a seeded database (`SEED_DB=true`).

### Companies

| ID                   | Name               | Type          |
| -------------------- | ------------------ | ------------- |
| `c-supplier-01`      | ChipCo Taiwan      | supplier      |
| `c-supplier-02`      | SteelWorks Germany | supplier      |
| `c-supplier-03`      | ChipCo Vietnam     | supplier      |
| `c-manufacturer-01`  | TechAssembly China | manufacturer  |
| `c-distributor-01`   | Euro Logistics     | distributor   |
| `c-retailer-01`      | TechShop EU        | retailer      |
| `c-customer-01`      | Acme Corp          | customer      |

### Products

| ID         | Name           | SKU          |
| ---------- | -------------- | ------------ |
| `prod-001` | Laptop Model X | LAP-X-001   |
| `prod-002` | Smartphone Pro | PHN-PRO-001 |

### Components

| ID         | Name          | Criticality |
| ---------- | ------------- | ----------- |
| `comp-001` | CPU Chip      | high        |
| `comp-002` | RAM Module    | medium      |
| `comp-003` | Steel Frame   | low         |
| `comp-004` | Display Panel | high        |
| `comp-005` | Battery Cell  | medium      |

### Locations

| ID          | Name               | Type                |
| ----------- | ------------------ | ------------------- |
| `loc-tw-01` | Taiwan Fab         | warehouse           |
| `loc-cn-01` | Shenzhen Assembly  | warehouse           |
| `loc-de-01` | Frankfurt Hub      | distribution_center |
| `loc-cz-01` | Prague Warehouse   | warehouse           |
| `loc-sg-01` | Singapore Port     | port                |

### Orders

| ID          | Product        | Status     |
| ----------- | -------------- | ---------- |
| `order-001` | Laptop Model X | in_transit |
| `order-002` | Smartphone Pro | pending    |
| `order-003` | Laptop Model X | delivered  |
| `order-004` | Smartphone Pro | delayed    |

### Routes

| ID          | Name                    |
| ----------- | ----------------------- |
| `route-001` | Taiwan-Singapore Sea    |
| `route-002` | Singapore-Frankfurt Air |
| `route-003` | Frankfurt-Prague Road   |
| `route-004` | Taiwan-Shenzhen Road    |

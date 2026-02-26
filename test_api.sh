#!/usr/bin/env bash
# ============================================================
# Supply Chain API — Endpoint Test Script
# ============================================================
# Prerequisites:
#   - Backend running on localhost:8080
#   - Database seeded (SEED_DB=true)
#   - curl and jq installed
#
# Usage:
#   chmod +x test_api.sh
#   ./test_api.sh              # run all tests
#   ./test_api.sh products     # run only a specific section
# ============================================================

set -euo pipefail

BASE="http://localhost:8080"
FILTER="${1:-all}"
PASS=0
FAIL=0

# Colors
GREEN='\033[0;32m'
RED='\033[0;31m'
CYAN='\033[0;36m'
YELLOW='\033[1;33m'
NC='\033[0m'

# ---- Helpers ----

separator() {
  echo ""
  echo -e "${CYAN}══════════════════════════════════════════════════════════${NC}"
  echo -e "${CYAN}  $1${NC}"
  echo -e "${CYAN}══════════════════════════════════════════════════════════${NC}"
}

run_test() {
  local description="$1"
  local method="$2"
  local url="$3"
  local data="${4:-}"
  local expected_code="${5:-200}"

  echo ""
  echo -e "${YELLOW}▸ ${description}${NC}"
  echo "  ${method} ${url}"

  if [[ -n "$data" ]]; then
    response=$(curl -s -w "\n%{http_code}" -X "$method" "$url" \
      -H "Content-Type: application/json" \
      -d "$data" 2>&1)
  else
    response=$(curl -s -w "\n%{http_code}" -X "$method" "$url" 2>&1)
  fi

  http_code=$(echo "$response" | tail -n1)
  body=$(echo "$response" | sed '$d')

  if [[ "$http_code" == "$expected_code" ]]; then
    echo -e "  ${GREEN}✓ Status: ${http_code}${NC}"
    PASS=$((PASS + 1))
  else
    echo -e "  ${RED}✗ Expected ${expected_code}, got ${http_code}${NC}"
    FAIL=$((FAIL + 1))
  fi

  # Pretty-print JSON body (truncated)
  if [[ -n "$body" ]] && command -v jq &>/dev/null; then
    echo "$body" | jq '.' 2>/dev/null | head -30
    local lines
    lines=$(echo "$body" | jq '.' 2>/dev/null | wc -l)
    if (( lines > 30 )); then
      echo "  ... (truncated, ${lines} lines total)"
    fi
  elif [[ -n "$body" ]]; then
    echo "$body" | head -10
  fi
}

should_run() {
  [[ "$FILTER" == "all" || "$FILTER" == "$1" ]]
}

# ============================================================
#  PRODUCTS
# ============================================================
if should_run "products"; then
  separator "PRODUCTS CRUD"

  run_test "List all products" \
    GET "${BASE}/api/products"

  run_test "Get single product by ID" \
    GET "${BASE}/api/products?id=prod-001"

  run_test "Get non-existent product (expect 404)" \
    GET "${BASE}/api/products?id=does-not-exist" "" 404

  run_test "Create a new product" \
    POST "${BASE}/api/products" \
    '{"name":"Test Widget","sku":"TST-001","price":19.99,"weight":0.3,"leadTime":5,"status":"active"}' \
    201

  run_test "Update product prod-001" \
    PUT "${BASE}/api/products?id=prod-001" \
    '{"name":"Laptop Model X (Updated)","sku":"LAP-X-001","price":1099.99,"weight":2.1,"leadTime":14,"status":"active"}'

  # NOTE: delete is destructive — creates a throwaway then deletes it
  run_test "Create throwaway product for delete test" \
    POST "${BASE}/api/products" \
    '{"name":"DeleteMe","sku":"DEL-001","price":1.0,"weight":0.1,"leadTime":1,"status":"active"}' \
    201
fi

# ============================================================
#  PRODUCTS — BOM
# ============================================================
if should_run "bom"; then
  separator "PRODUCTS — BOM"

  run_test "Get BOM for prod-001" \
    GET "${BASE}/api/products/bom?id=prod-001"

  run_test "Get detailed BOM for prod-001" \
    GET "${BASE}/api/products/bom/detailed?id=prod-001"

  run_test "Get alternative suppliers for prod-001" \
    GET "${BASE}/api/products/alternative-suppliers?id=prod-001"

  run_test "Add component to BOM (prod-002 + comp-003)" \
    POST "${BASE}/api/products/bom?id=prod-002" \
    '{"componentId":"comp-003","quantity":1,"position":4}' \
    201

  run_test "Update BOM component quantity (prod-001, comp-001)" \
    PUT "${BASE}/api/products/bom?id=prod-001&componentId=comp-001" \
    '{"quantity":3}'
fi

# ============================================================
#  COMPANIES
# ============================================================
if should_run "companies"; then
  separator "COMPANIES CRUD"

  run_test "List all companies" \
    GET "${BASE}/api/companies"

  run_test "Get single company" \
    GET "${BASE}/api/companies?id=c-supplier-01"

  run_test "Create a new company" \
    POST "${BASE}/api/companies" \
    '{"name":"Test Supplier","type":"supplier","country":"Japan","coordinates":{"lat":35.689,"lng":139.692},"reliability":0.9}' \
    201

  run_test "Update company c-supplier-02" \
    PUT "${BASE}/api/companies?id=c-supplier-02" \
    '{"name":"SteelWorks Germany (Updated)","type":"supplier","country":"Germany","coordinates":{"lat":50.110,"lng":8.682},"reliability":0.94}'
fi

# ============================================================
#  COMPANIES — RISK ASSESSMENT
# ============================================================
if should_run "risk"; then
  separator "RISK ASSESSMENT"

  run_test "Risk assessment for c-supplier-01" \
    GET "${BASE}/api/companies/risk-assessment?id=c-supplier-01"

  run_test "Risk assessment for c-supplier-03 (lower reliability)" \
    GET "${BASE}/api/companies/risk-assessment?id=c-supplier-03"
fi

# ============================================================
#  COMPONENTS
# ============================================================
if should_run "components"; then
  separator "COMPONENTS CRUD"

  run_test "List all components" \
    GET "${BASE}/api/components"

  run_test "Get single component" \
    GET "${BASE}/api/components?id=comp-001"

  run_test "Create a component" \
    POST "${BASE}/api/components" \
    '{"name":"Test Resistor","price":0.05,"quantity":10000,"criticality":"low"}' \
    201

  run_test "Update component comp-002" \
    PUT "${BASE}/api/components?id=comp-002" \
    '{"name":"RAM Module (DDR5)","price":55.0,"quantity":500,"criticality":"medium"}'
fi

# ============================================================
#  ORDERS
# ============================================================
if should_run "orders"; then
  separator "ORDERS"

  run_test "List all orders" \
    GET "${BASE}/api/orders"

  run_test "Get single order" \
    GET "${BASE}/api/orders?id=order-001"

  run_test "Create a new order" \
    POST "${BASE}/api/orders" \
    '{"orderDate":"2024-04-01","dueDate":"2024-04-20","quantity":25,"status":"pending","cost":12500.0,"productId":"prod-002","productQuantity":25,"unitPrice":350.0,"customerId":"c-customer-01","supplierId":"c-distributor-01"}' \
    201

  run_test "Update order status (order-002 → in_transit)" \
    PUT "${BASE}/api/orders/status?id=order-002" \
    '{"status":"in_transit"}'
fi

# ============================================================
#  ORDERS — SUPPLY PATH & COST
# ============================================================
if should_run "supply"; then
  separator "SUPPLY PATH & COST BREAKDOWN"

  run_test "Supply path for order-001" \
    GET "${BASE}/api/orders/supply-path?orderId=order-001"

  run_test "Cost breakdown for order-001" \
    GET "${BASE}/api/orders/cost-breakdown?orderId=order-001"

  run_test "Cost breakdown for order-003" \
    GET "${BASE}/api/orders/cost-breakdown?orderId=order-003"
fi

# ============================================================
#  LOCATIONS
# ============================================================
if should_run "locations"; then
  separator "LOCATIONS"

  run_test "List all locations" \
    GET "${BASE}/api/locations"

  run_test "Get single location" \
    GET "${BASE}/api/locations?id=loc-de-01"

  run_test "Create a location" \
    POST "${BASE}/api/locations" \
    '{"name":"Test Depot","type":"warehouse","coordinates":{"lat":48.856,"lng":2.352},"capacity":1000}' \
    201

  run_test "Inventory status for loc-de-01 (Frankfurt Hub)" \
    GET "${BASE}/api/locations/inventory-status?id=loc-de-01"

  run_test "Inventory status for loc-cz-01 (Prague Warehouse)" \
    GET "${BASE}/api/locations/inventory-status?id=loc-cz-01"
fi

# ============================================================
#  ROUTES
# ============================================================
if should_run "routes"; then
  separator "OPTIMAL ROUTES"

  run_test "Optimal route: Taiwan → Prague" \
    GET "${BASE}/api/routes/optimal?from=loc-tw-01&to=loc-cz-01"

  run_test "Optimal route: Taiwan → Shenzhen" \
    GET "${BASE}/api/routes/optimal?from=loc-tw-01&to=loc-cn-01"

  run_test "Optimal route: no path (expect 404)" \
    GET "${BASE}/api/routes/optimal?from=loc-cz-01&to=loc-tw-01" "" 404
fi

# ============================================================
#  ANALYTICS
# ============================================================
if should_run "analytics"; then
  separator "ANALYTICS"

  run_test "Supply chain health" \
    GET "${BASE}/api/analytics/supply-chain-health"

  run_test "Impact analysis for c-supplier-01" \
    GET "${BASE}/api/analytics/impact-analysis?supplier=c-supplier-01"

  run_test "Impact analysis for c-supplier-03" \
    GET "${BASE}/api/analytics/impact-analysis?supplier=c-supplier-03"
fi

# ============================================================
#  PREDICTIONS
# ============================================================
if should_run "predictions"; then
  separator "PREDICTIONS"

  run_test "Forecast delays (default 3 months)" \
    GET "${BASE}/api/analytics/forecast-delays"

  run_test "Forecast delays (6 months)" \
    GET "${BASE}/api/analytics/forecast-delays?months=6"

  run_test "Stock level forecast for prod-001 (default 6 months)" \
    GET "${BASE}/api/analytics/stock-levels?product=prod-001"

  run_test "Stock level forecast for prod-002 (12 months)" \
    GET "${BASE}/api/analytics/stock-levels?product=prod-002&months=12"
fi

# ============================================================
#  SUMMARY
# ============================================================
echo ""
echo -e "${CYAN}══════════════════════════════════════════════════════════${NC}"
echo -e "${CYAN}  RESULTS${NC}"
echo -e "${CYAN}══════════════════════════════════════════════════════════${NC}"
TOTAL=$((PASS + FAIL))
echo -e "  Total: ${TOTAL}   ${GREEN}Passed: ${PASS}${NC}   ${RED}Failed: ${FAIL}${NC}"
echo ""

if (( FAIL > 0 )); then
  exit 1
fi

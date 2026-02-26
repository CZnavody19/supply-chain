package db

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
	"go.uber.org/zap"
)

func (ds *DatabaseStore) SeedDatabase(ctx context.Context) error {
	session := ds.newSession(ctx)
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		// ---- Companies ----
		_, err := tx.Run(ctx, `
			MERGE (c:Company {id: 'c-supplier-01'})
			SET c.name = 'ChipCo Taiwan', c.type = 'supplier', c.country = 'Taiwan',
			    c.lat = 25.033, c.lng = 121.565, c.reliability = 0.95

			MERGE (c2:Company {id: 'c-supplier-02'})
			SET c2.name = 'SteelWorks Germany', c2.type = 'supplier', c2.country = 'Germany',
			    c2.lat = 50.110, c2.lng = 8.682, c2.reliability = 0.92

			MERGE (c3:Company {id: 'c-supplier-03'})
			SET c3.name = 'ChipCo Vietnam', c3.type = 'supplier', c3.country = 'Vietnam',
			    c3.lat = 10.823, c3.lng = 106.630, c3.reliability = 0.85

			MERGE (c4:Company {id: 'c-manufacturer-01'})
			SET c4.name = 'TechAssembly China', c4.type = 'manufacturer', c4.country = 'China',
			    c4.lat = 22.543, c4.lng = 114.058, c4.reliability = 0.93

			MERGE (c5:Company {id: 'c-distributor-01'})
			SET c5.name = 'Euro Logistics', c5.type = 'distributor', c5.country = 'Germany',
			    c5.lat = 50.937, c5.lng = 6.960, c5.reliability = 0.96

			MERGE (c6:Company {id: 'c-retailer-01'})
			SET c6.name = 'TechShop EU', c6.type = 'retailer', c6.country = 'Czech Republic',
			    c6.lat = 50.075, c6.lng = 14.437, c6.reliability = 0.97

			MERGE (c7:Company {id: 'c-customer-01'})
			SET c7.name = 'Acme Corp', c7.type = 'customer', c7.country = 'Czech Republic',
			    c7.lat = 49.195, c7.lng = 16.608, c7.reliability = 1.0
		`, nil)
		if err != nil {
			return nil, err
		}

		// ---- Products ----
		_, err = tx.Run(ctx, `
			MERGE (p1:Product {id: 'prod-001'})
			SET p1.name = 'Laptop Model X', p1.sku = 'LAP-X-001', p1.price = 999.99,
			    p1.weight = 2.1, p1.leadTime = 14, p1.status = 'active'

			MERGE (p2:Product {id: 'prod-002'})
			SET p2.name = 'Smartphone Pro', p2.sku = 'PHN-PRO-001', p2.price = 699.99,
			    p2.weight = 0.2, p2.leadTime = 7, p2.status = 'active'
		`, nil)
		if err != nil {
			return nil, err
		}

		// ---- Components ----
		_, err = tx.Run(ctx, `
			MERGE (c1:Component {id: 'comp-001'})
			SET c1.name = 'CPU Chip', c1.price = 150.0, c1.quantity = 0, c1.criticality = 'high'

			MERGE (c2:Component {id: 'comp-002'})
			SET c2.name = 'RAM Module', c2.price = 45.0, c2.quantity = 0, c2.criticality = 'medium'

			MERGE (c3:Component {id: 'comp-003'})
			SET c3.name = 'Steel Frame', c3.price = 20.0, c3.quantity = 0, c3.criticality = 'low'

			MERGE (c4:Component {id: 'comp-004'})
			SET c4.name = 'Display Panel', c4.price = 85.0, c4.quantity = 0, c4.criticality = 'high'

			MERGE (c5:Component {id: 'comp-005'})
			SET c5.name = 'Battery Cell', c5.price = 35.0, c5.quantity = 0, c5.criticality = 'medium'
		`, nil)
		if err != nil {
			return nil, err
		}

		// ---- Locations ----
		_, err = tx.Run(ctx, `
			MERGE (l1:Location {id: 'loc-tw-01'})
			SET l1.name = 'Taiwan Fab', l1.type = 'warehouse', l1.lat = 25.033, l1.lng = 121.565, l1.capacity = 5000

			MERGE (l2:Location {id: 'loc-cn-01'})
			SET l2.name = 'Shenzhen Assembly', l2.type = 'warehouse', l2.lat = 22.543, l2.lng = 114.058, l2.capacity = 10000

			MERGE (l3:Location {id: 'loc-de-01'})
			SET l3.name = 'Frankfurt Hub', l3.type = 'distribution_center', l3.lat = 50.110, l3.lng = 8.682, l3.capacity = 8000

			MERGE (l4:Location {id: 'loc-cz-01'})
			SET l4.name = 'Prague Warehouse', l4.type = 'warehouse', l4.lat = 50.075, l4.lng = 14.437, l4.capacity = 3000

			MERGE (l5:Location {id: 'loc-sg-01'})
			SET l5.name = 'Singapore Port', l5.type = 'port', l5.lat = 1.352, l5.lng = 103.820, l5.capacity = 20000
		`, nil)
		if err != nil {
			return nil, err
		}

		// ---- Routes ----
		_, err = tx.Run(ctx, `
			MERGE (r1:Route {id: 'route-001'})
			SET r1.name = 'Taiwan-Singapore Sea', r1.distance = 2500, r1.estimatedTime = 120, r1.cost = 3000, r1.reliability = 0.92

			MERGE (r2:Route {id: 'route-002'})
			SET r2.name = 'Singapore-Frankfurt Air', r2.distance = 10000, r2.estimatedTime = 14, r2.cost = 8000, r2.reliability = 0.95

			MERGE (r3:Route {id: 'route-003'})
			SET r3.name = 'Frankfurt-Prague Road', r3.distance = 350, r3.estimatedTime = 5, r3.cost = 500, r3.reliability = 0.98

			MERGE (r4:Route {id: 'route-004'})
			SET r4.name = 'Taiwan-Shenzhen Road', r4.distance = 800, r4.estimatedTime = 12, r4.cost = 400, r4.reliability = 0.97
		`, nil)
		if err != nil {
			return nil, err
		}

		// ---- COMPOSED_OF (BOM) ----
		_, err = tx.Run(ctx, `
			MATCH (p1:Product {id: 'prod-001'}), (c1:Component {id: 'comp-001'})
			MERGE (p1)-[:COMPOSED_OF {quantity: 1, position: 1}]->(c1)

			WITH 1 AS dummy
			MATCH (p1:Product {id: 'prod-001'}), (c2:Component {id: 'comp-002'})
			MERGE (p1)-[:COMPOSED_OF {quantity: 2, position: 2}]->(c2)

			WITH 1 AS dummy
			MATCH (p1:Product {id: 'prod-001'}), (c3:Component {id: 'comp-003'})
			MERGE (p1)-[:COMPOSED_OF {quantity: 1, position: 3}]->(c3)

			WITH 1 AS dummy
			MATCH (p2:Product {id: 'prod-002'}), (c4:Component {id: 'comp-004'})
			MERGE (p2)-[:COMPOSED_OF {quantity: 1, position: 1}]->(c4)

			WITH 1 AS dummy
			MATCH (p2:Product {id: 'prod-002'}), (c5:Component {id: 'comp-005'})
			MERGE (p2)-[:COMPOSED_OF {quantity: 1, position: 2}]->(c5)

			WITH 1 AS dummy
			MATCH (p2:Product {id: 'prod-002'}), (c1:Component {id: 'comp-001'})
			MERGE (p2)-[:COMPOSED_OF {quantity: 1, position: 3}]->(c1)
		`, nil)
		if err != nil {
			return nil, err
		}

		// ---- SUPPLIED_BY ----
		_, err = tx.Run(ctx, `
			MATCH (c1:Component {id: 'comp-001'}), (s1:Company {id: 'c-supplier-01'})
			MERGE (c1)-[:SUPPLIED_BY {price: 120.0, leadTime: 10, minOrder: 100}]->(s1)

			WITH 1 AS dummy
			MATCH (c1:Component {id: 'comp-001'}), (s3:Company {id: 'c-supplier-03'})
			MERGE (c1)-[:SUPPLIED_BY {price: 135.0, leadTime: 14, minOrder: 50}]->(s3)

			WITH 1 AS dummy
			MATCH (c2:Component {id: 'comp-002'}), (s1:Company {id: 'c-supplier-01'})
			MERGE (c2)-[:SUPPLIED_BY {price: 40.0, leadTime: 7, minOrder: 200}]->(s1)

			WITH 1 AS dummy
			MATCH (c3:Component {id: 'comp-003'}), (s2:Company {id: 'c-supplier-02'})
			MERGE (c3)-[:SUPPLIED_BY {price: 15.0, leadTime: 5, minOrder: 500}]->(s2)

			WITH 1 AS dummy
			MATCH (c4:Component {id: 'comp-004'}), (s3:Company {id: 'c-supplier-03'})
			MERGE (c4)-[:SUPPLIED_BY {price: 80.0, leadTime: 12, minOrder: 100}]->(s3)

			WITH 1 AS dummy
			MATCH (c5:Component {id: 'comp-005'}), (s3:Company {id: 'c-supplier-03'})
			MERGE (c5)-[:SUPPLIED_BY {price: 30.0, leadTime: 8, minOrder: 200}]->(s3)
		`, nil)
		if err != nil {
			return nil, err
		}

		// ---- MANUFACTURES ----
		_, err = tx.Run(ctx, `
			MATCH (mfg:Company {id: 'c-manufacturer-01'}), (p1:Product {id: 'prod-001'})
			MERGE (mfg)-[:MANUFACTURES {capacity: 1000, unitCost: 450.0, qualityScore: 0.95}]->(p1)

			WITH 1 AS dummy
			MATCH (mfg:Company {id: 'c-manufacturer-01'}), (p2:Product {id: 'prod-002'})
			MERGE (mfg)-[:MANUFACTURES {capacity: 2000, unitCost: 280.0, qualityScore: 0.93}]->(p2)
		`, nil)
		if err != nil {
			return nil, err
		}

		// ---- DISTRIBUTOR_OF ----
		_, err = tx.Run(ctx, `
			MATCH (d:Company {id: 'c-distributor-01'}), (p1:Product {id: 'prod-001'})
			MERGE (d)-[:DISTRIBUTOR_OF {stock: 500, lastRestocked: '2024-01-15'}]->(p1)

			WITH 1 AS dummy
			MATCH (d:Company {id: 'c-distributor-01'}), (p2:Product {id: 'prod-002'})
			MERGE (d)-[:DISTRIBUTOR_OF {stock: 300, lastRestocked: '2024-01-20'}]->(p2)
		`, nil)
		if err != nil {
			return nil, err
		}

		// ---- LOCATED_AT ----
		_, err = tx.Run(ctx, `
			MATCH (c:Company {id: 'c-supplier-01'}), (l:Location {id: 'loc-tw-01'})
			MERGE (c)-[:LOCATED_AT {since: '2018-01-01'}]->(l)

			WITH 1 AS dummy
			MATCH (c:Company {id: 'c-supplier-02'}), (l:Location {id: 'loc-de-01'})
			MERGE (c)-[:LOCATED_AT {since: '2015-06-01'}]->(l)

			WITH 1 AS dummy
			MATCH (c:Company {id: 'c-supplier-03'}), (l:Location {id: 'loc-sg-01'})
			MERGE (c)-[:LOCATED_AT {since: '2019-03-01'}]->(l)

			WITH 1 AS dummy
			MATCH (c:Company {id: 'c-manufacturer-01'}), (l:Location {id: 'loc-cn-01'})
			MERGE (c)-[:LOCATED_AT {since: '2016-01-01'}]->(l)

			WITH 1 AS dummy
			MATCH (c:Company {id: 'c-distributor-01'}), (l:Location {id: 'loc-de-01'})
			MERGE (c)-[:LOCATED_AT {since: '2017-09-01'}]->(l)

			WITH 1 AS dummy
			MATCH (c:Company {id: 'c-retailer-01'}), (l:Location {id: 'loc-cz-01'})
			MERGE (c)-[:LOCATED_AT {since: '2020-01-01'}]->(l)

			WITH 1 AS dummy
			MATCH (c:Company {id: 'c-customer-01'}), (l:Location {id: 'loc-cz-01'})
			MERGE (c)-[:LOCATED_AT {since: '2021-05-01'}]->(l)
		`, nil)
		if err != nil {
			return nil, err
		}

		// ---- CONNECTED_TO (location connections with route data) ----
		_, err = tx.Run(ctx, `
			MATCH (l1:Location {id: 'loc-tw-01'}), (l2:Location {id: 'loc-sg-01'})
			MERGE (l1)-[:CONNECTED_TO {distance: 2500, time: 120, cost: 3000, routeId: 'route-001'}]->(l2)

			WITH 1 AS dummy
			MATCH (l1:Location {id: 'loc-sg-01'}), (l2:Location {id: 'loc-de-01'})
			MERGE (l1)-[:CONNECTED_TO {distance: 10000, time: 14, cost: 8000, routeId: 'route-002'}]->(l2)

			WITH 1 AS dummy
			MATCH (l1:Location {id: 'loc-de-01'}), (l2:Location {id: 'loc-cz-01'})
			MERGE (l1)-[:CONNECTED_TO {distance: 350, time: 5, cost: 500, routeId: 'route-003'}]->(l2)

			WITH 1 AS dummy
			MATCH (l1:Location {id: 'loc-tw-01'}), (l2:Location {id: 'loc-cn-01'})
			MERGE (l1)-[:CONNECTED_TO {distance: 800, time: 12, cost: 400, routeId: 'route-004'}]->(l2)

			WITH 1 AS dummy
			MATCH (l1:Location {id: 'loc-cn-01'}), (l2:Location {id: 'loc-sg-01'})
			MERGE (l1)-[:CONNECTED_TO {distance: 2800, time: 96, cost: 2500, routeId: 'route-001'}]->(l2)
		`, nil)
		if err != nil {
			return nil, err
		}

		// ---- STORED_AT ----
		_, err = tx.Run(ctx, `
			MATCH (p:Product {id: 'prod-001'}), (l:Location {id: 'loc-de-01'})
			MERGE (p)-[:STORED_AT {quantity: 500, lastRestockDate: '2024-01-15'}]->(l)

			WITH 1 AS dummy
			MATCH (p:Product {id: 'prod-001'}), (l:Location {id: 'loc-cz-01'})
			MERGE (p)-[:STORED_AT {quantity: 200, lastRestockDate: '2024-01-10'}]->(l)

			WITH 1 AS dummy
			MATCH (p:Product {id: 'prod-002'}), (l:Location {id: 'loc-de-01'})
			MERGE (p)-[:STORED_AT {quantity: 300, lastRestockDate: '2024-01-20'}]->(l)
		`, nil)
		if err != nil {
			return nil, err
		}

		// ---- SUPPLIES (B2B) ----
		_, err = tx.Run(ctx, `
			MATCH (s:Company {id: 'c-supplier-01'}), (m:Company {id: 'c-manufacturer-01'})
			MERGE (s)-[:SUPPLIES {contractSince: '2020-01-01', minOrder: 100, leadTime: 10}]->(m)

			WITH 1 AS dummy
			MATCH (s:Company {id: 'c-supplier-03'}), (m:Company {id: 'c-manufacturer-01'})
			MERGE (s)-[:SUPPLIES {contractSince: '2021-06-01', minOrder: 50, leadTime: 14}]->(m)

			WITH 1 AS dummy
			MATCH (m:Company {id: 'c-manufacturer-01'}), (d:Company {id: 'c-distributor-01'})
			MERGE (m)-[:SUPPLIES {contractSince: '2019-01-01', minOrder: 50, leadTime: 7}]->(d)

			WITH 1 AS dummy
			MATCH (d:Company {id: 'c-distributor-01'}), (r:Company {id: 'c-retailer-01'})
			MERGE (d)-[:SUPPLIES {contractSince: '2020-03-01', minOrder: 10, leadTime: 3}]->(r)
		`, nil)
		if err != nil {
			return nil, err
		}

		// ---- Orders ----
		_, err = tx.Run(ctx, `
			MERGE (o1:Order {id: 'order-001'})
			SET o1.orderDate = '2024-02-01', o1.dueDate = '2024-02-28', o1.quantity = 100,
			    o1.status = 'in_transit', o1.cost = 45000.0

			MERGE (o2:Order {id: 'order-002'})
			SET o2.orderDate = '2024-02-05', o2.dueDate = '2024-03-05', o2.quantity = 50,
			    o2.status = 'pending', o2.cost = 17500.0

			MERGE (o3:Order {id: 'order-003'})
			SET o3.orderDate = '2024-01-10', o3.dueDate = '2024-01-30', o3.quantity = 200,
			    o3.status = 'delivered', o3.cost = 88000.0

			MERGE (o4:Order {id: 'order-004'})
			SET o4.orderDate = '2024-01-15', o4.dueDate = '2024-02-01', o4.quantity = 75,
			    o4.status = 'delayed', o4.cost = 26250.0
		`, nil)
		if err != nil {
			return nil, err
		}

		// ---- Order relationships ----
		_, err = tx.Run(ctx, `
			MATCH (o:Order {id: 'order-001'}), (p:Product {id: 'prod-001'})
			MERGE (o)-[:CONTAINS {quantity: 100, unitPrice: 450.0}]->(p)
			WITH o
			MATCH (cust:Company {id: 'c-customer-01'})
			MERGE (o)-[:FROM]->(cust)
			WITH o
			MATCH (sup:Company {id: 'c-manufacturer-01'})
			MERGE (o)-[:PLACED_WITH]->(sup)
		`, nil)
		if err != nil {
			return nil, err
		}

		_, err = tx.Run(ctx, `
			MATCH (o:Order {id: 'order-002'}), (p:Product {id: 'prod-002'})
			MERGE (o)-[:CONTAINS {quantity: 50, unitPrice: 350.0}]->(p)
			WITH o
			MATCH (cust:Company {id: 'c-customer-01'})
			MERGE (o)-[:FROM]->(cust)
			WITH o
			MATCH (sup:Company {id: 'c-distributor-01'})
			MERGE (o)-[:PLACED_WITH]->(sup)
		`, nil)
		if err != nil {
			return nil, err
		}

		_, err = tx.Run(ctx, `
			MATCH (o:Order {id: 'order-003'}), (p:Product {id: 'prod-001'})
			MERGE (o)-[:CONTAINS {quantity: 200, unitPrice: 440.0}]->(p)
			WITH o
			MATCH (cust:Company {id: 'c-retailer-01'})
			MERGE (o)-[:FROM]->(cust)
			WITH o
			MATCH (sup:Company {id: 'c-manufacturer-01'})
			MERGE (o)-[:PLACED_WITH]->(sup)
		`, nil)
		if err != nil {
			return nil, err
		}

		_, err = tx.Run(ctx, `
			MATCH (o:Order {id: 'order-004'}), (p:Product {id: 'prod-002'})
			MERGE (o)-[:CONTAINS {quantity: 75, unitPrice: 350.0}]->(p)
			WITH o
			MATCH (cust:Company {id: 'c-retailer-01'})
			MERGE (o)-[:FROM]->(cust)
			WITH o
			MATCH (sup:Company {id: 'c-distributor-01'})
			MERGE (o)-[:PLACED_WITH]->(sup)
		`, nil)
		if err != nil {
			return nil, err
		}

		// ---- SHIPPED_VIA ----
		_, err = tx.Run(ctx, `
			MATCH (o:Order {id: 'order-001'}), (r:Route {id: 'route-002'})
			MERGE (o)-[:SHIPPED_VIA {departureDate: '2024-02-15', arrivalDate: '2024-02-16'}]->(r)

			WITH 1 AS dummy
			MATCH (o:Order {id: 'order-003'}), (r:Route {id: 'route-003'})
			MERGE (o)-[:SHIPPED_VIA {departureDate: '2024-01-25', arrivalDate: '2024-01-26'}]->(r)
		`, nil)
		if err != nil {
			return nil, err
		}

		return nil, nil
	})

	if err != nil {
		zap.L().Error("Error seeding database", zap.Error(err))
		return err
	}
	zap.L().Info("Database seeded successfully")
	return nil
}

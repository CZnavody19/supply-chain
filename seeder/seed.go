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
		var err error
		_, err = tx.Run(ctx, `
			MERGE (c1:Company {id: 'c-001'})
			SET c1.name = 'Company 1', c1.type = 'customer', c1.country = 'Germany', c1.lat = 51.161, c1.lng = 9.807, c1.reliability = 0.89

			MERGE (c2:Company {id: 'c-002'})
			SET c2.name = 'Company 2', c2.type = 'supplier', c2.country = 'Germany', c2.lat = 51.535, c2.lng = 9.844, c2.reliability = 0.83

			MERGE (c3:Company {id: 'c-003'})
			SET c3.name = 'Company 3', c3.type = 'manufacturer', c3.country = 'Vietnam', c3.lat = 13.143, c3.lng = 108.871, c3.reliability = 0.72

			MERGE (c4:Company {id: 'c-004'})
			SET c4.name = 'Company 4', c4.type = 'retailer', c4.country = 'Vietnam', c4.lat = 13.556, c4.lng = 107.412, c4.reliability = 0.90

			MERGE (c5:Company {id: 'c-005'})
			SET c5.name = 'Company 5', c5.type = 'supplier', c5.country = 'Czech Republic', c5.lat = 50.639, c5.lng = 14.719, c5.reliability = 0.76

			MERGE (c6:Company {id: 'c-006'})
			SET c6.name = 'Company 6', c6.type = 'customer', c6.country = 'Vietnam', c6.lat = 13.863, c6.lng = 108.503, c6.reliability = 0.75

			MERGE (c7:Company {id: 'c-007'})
			SET c7.name = 'Company 7', c7.type = 'supplier', c7.country = 'Germany', c7.lat = 51.655, c7.lng = 9.526, c7.reliability = 0.84

			MERGE (c8:Company {id: 'c-008'})
			SET c8.name = 'Company 8', c8.type = 'retailer', c8.country = 'Vietnam', c8.lat = 13.283, c8.lng = 108.324, c8.reliability = 0.73

			MERGE (c9:Company {id: 'c-009'})
			SET c9.name = 'Company 9', c9.type = 'supplier', c9.country = 'Czech Republic', c9.lat = 49.449, c9.lng = 15.746, c9.reliability = 0.96

			MERGE (c10:Company {id: 'c-010'})
			SET c10.name = 'Company 10', c10.type = 'manufacturer', c10.country = 'Vietnam', c10.lat = 13.007, c10.lng = 108.024, c10.reliability = 0.81

			MERGE (c11:Company {id: 'c-011'})
			SET c11.name = 'Company 11', c11.type = 'supplier', c11.country = 'Taiwan', c11.lat = 24.840, c11.lng = 121.163, c11.reliability = 0.99

			MERGE (c12:Company {id: 'c-012'})
			SET c12.name = 'Company 12', c12.type = 'supplier', c12.country = 'Vietnam', c12.lat = 14.217, c12.lng = 108.281, c12.reliability = 0.91

			MERGE (c13:Company {id: 'c-013'})
			SET c13.name = 'Company 13', c13.type = 'retailer', c13.country = 'USA', c13.lat = 37.694, c13.lng = -94.868, c13.reliability = 0.99

			MERGE (c14:Company {id: 'c-014'})
			SET c14.name = 'Company 14', c14.type = 'manufacturer', c14.country = 'Germany', c14.lat = 51.801, c14.lng = 9.439, c14.reliability = 0.89

			MERGE (c15:Company {id: 'c-015'})
			SET c15.name = 'Company 15', c15.type = 'retailer', c15.country = 'Japan', c15.lat = 36.893, c15.lng = 138.996, c15.reliability = 0.97

			MERGE (c16:Company {id: 'c-016'})
			SET c16.name = 'Company 16', c16.type = 'supplier', c16.country = 'Germany', c16.lat = 52.097, c16.lng = 10.994, c16.reliability = 0.90

			MERGE (c17:Company {id: 'c-017'})
			SET c17.name = 'Company 17', c17.type = 'customer', c17.country = 'USA', c17.lat = 37.334, c17.lng = -95.038, c17.reliability = 0.78

			MERGE (c18:Company {id: 'c-018'})
			SET c18.name = 'Company 18', c18.type = 'supplier', c18.country = 'Vietnam', c18.lat = 13.582, c18.lng = 107.415, c18.reliability = 0.90

			MERGE (c19:Company {id: 'c-019'})
			SET c19.name = 'Company 19', c19.type = 'manufacturer', c19.country = 'Germany', c19.lat = 50.607, c19.lng = 9.844, c19.reliability = 0.87

			MERGE (c20:Company {id: 'c-020'})
			SET c20.name = 'Company 20', c20.type = 'supplier', c20.country = 'Germany', c20.lat = 52.001, c20.lng = 10.945, c20.reliability = 0.88

			MERGE (c21:Company {id: 'c-021'})
			SET c21.name = 'Company 21', c21.type = 'supplier', c21.country = 'Taiwan', c21.lat = 24.704, c21.lng = 122.390, c21.reliability = 0.95

			MERGE (c22:Company {id: 'c-022'})
			SET c22.name = 'Company 22', c22.type = 'supplier', c22.country = 'Vietnam', c22.lat = 14.415, c22.lng = 107.920, c22.reliability = 0.90

			MERGE (c23:Company {id: 'c-023'})
			SET c23.name = 'Company 23', c23.type = 'manufacturer', c23.country = 'Japan', c23.lat = 37.052, c23.lng = 137.968, c23.reliability = 0.99

			MERGE (c24:Company {id: 'c-024'})
			SET c24.name = 'Company 24', c24.type = 'supplier', c24.country = 'Czech Republic', c24.lat = 49.688, c24.lng = 14.529, c24.reliability = 0.77

			MERGE (c25:Company {id: 'c-025'})
			SET c25.name = 'Company 25', c25.type = 'supplier', c25.country = 'China', c25.lat = 36.525, c25.lng = 104.311, c25.reliability = 0.82

			MERGE (c26:Company {id: 'c-026'})
			SET c26.name = 'Company 26', c26.type = 'distributor', c26.country = 'Czech Republic', c26.lat = 50.757, c26.lng = 15.866, c26.reliability = 0.91

			MERGE (c27:Company {id: 'c-027'})
			SET c27.name = 'Company 27', c27.type = 'manufacturer', c27.country = 'China', c27.lat = 35.851, c27.lng = 103.819, c27.reliability = 0.90

			MERGE (c28:Company {id: 'c-028'})
			SET c28.name = 'Company 28', c28.type = 'customer', c28.country = 'China', c28.lat = 35.531, c28.lng = 103.798, c28.reliability = 0.79

			MERGE (c29:Company {id: 'c-029'})
			SET c29.name = 'Company 29', c29.type = 'customer', c29.country = 'Japan', c29.lat = 36.229, c29.lng = 138.391, c29.reliability = 0.95

			MERGE (c30:Company {id: 'c-030'})
			SET c30.name = 'Company 30', c30.type = 'supplier', c30.country = 'China', c30.lat = 35.987, c30.lng = 103.788, c30.reliability = 0.98
		`, nil)
		if err != nil { return nil, err }

		_, err = tx.Run(ctx, `
			MERGE (l1:Location {id: 'loc-001'})
			SET l1.name = 'Location 1', l1.type = 'port', l1.lat = 37.000, l1.lng = -95.700, l1.capacity = 5157

			MERGE (l2:Location {id: 'loc-002'})
			SET l2.name = 'Location 2', l2.type = 'warehouse', l2.lat = 51.100, l2.lng = 10.400, l2.capacity = 48050

			MERGE (l3:Location {id: 'loc-003'})
			SET l3.name = 'Location 3', l3.type = 'distribution_center', l3.lat = 37.000, l3.lng = -95.700, l3.capacity = 25136

			MERGE (l4:Location {id: 'loc-004'})
			SET l4.name = 'Location 4', l4.type = 'warehouse', l4.lat = 51.100, l4.lng = 10.400, l4.capacity = 13422

			MERGE (l5:Location {id: 'loc-005'})
			SET l5.name = 'Location 5', l5.type = 'port', l5.lat = 51.100, l5.lng = 10.400, l5.capacity = 4878

			MERGE (l6:Location {id: 'loc-006'})
			SET l6.name = 'Location 6', l6.type = 'port', l6.lat = 14.000, l6.lng = 108.200, l6.capacity = 1672

			MERGE (l7:Location {id: 'loc-007'})
			SET l7.name = 'Location 7', l7.type = 'distribution_center', l7.lat = 36.200, l7.lng = 138.200, l7.capacity = 13389

			MERGE (l8:Location {id: 'loc-008'})
			SET l8.name = 'Location 8', l8.type = 'warehouse', l8.lat = 37.000, l8.lng = -95.700, l8.capacity = 13574

			MERGE (l9:Location {id: 'loc-009'})
			SET l9.name = 'Location 9', l9.type = 'warehouse', l9.lat = 36.200, l9.lng = 138.200, l9.capacity = 40086

			MERGE (l10:Location {id: 'loc-010'})
			SET l10.name = 'Location 10', l10.type = 'distribution_center', l10.lat = 35.800, l10.lng = 104.100, l10.capacity = 49551

			MERGE (l11:Location {id: 'loc-011'})
			SET l11.name = 'Location 11', l11.type = 'port', l11.lat = 49.800, l11.lng = 15.400, l11.capacity = 45396

			MERGE (l12:Location {id: 'loc-012'})
			SET l12.name = 'Location 12', l12.type = 'distribution_center', l12.lat = 25.000, l12.lng = 121.500, l12.capacity = 10186

			MERGE (l13:Location {id: 'loc-013'})
			SET l13.name = 'Location 13', l13.type = 'port', l13.lat = 14.000, l13.lng = 108.200, l13.capacity = 4917

			MERGE (l14:Location {id: 'loc-014'})
			SET l14.name = 'Location 14', l14.type = 'warehouse', l14.lat = 25.000, l14.lng = 121.500, l14.capacity = 37631

			MERGE (l15:Location {id: 'loc-015'})
			SET l15.name = 'Location 15', l15.type = 'warehouse', l15.lat = 49.800, l15.lng = 15.400, l15.capacity = 31134

			MERGE (l16:Location {id: 'loc-016'})
			SET l16.name = 'Location 16', l16.type = 'distribution_center', l16.lat = 35.800, l16.lng = 104.100, l16.capacity = 39374

			MERGE (l17:Location {id: 'loc-017'})
			SET l17.name = 'Location 17', l17.type = 'warehouse', l17.lat = 51.100, l17.lng = 10.400, l17.capacity = 19014

			MERGE (l18:Location {id: 'loc-018'})
			SET l18.name = 'Location 18', l18.type = 'port', l18.lat = 25.000, l18.lng = 121.500, l18.capacity = 1128

			MERGE (l19:Location {id: 'loc-019'})
			SET l19.name = 'Location 19', l19.type = 'warehouse', l19.lat = 49.800, l19.lng = 15.400, l19.capacity = 7804

			MERGE (l20:Location {id: 'loc-020'})
			SET l20.name = 'Location 20', l20.type = 'warehouse', l20.lat = 36.200, l20.lng = 138.200, l20.capacity = 25226
		`, nil)
		if err != nil { return nil, err }

		_, err = tx.Run(ctx, `
			MERGE (p1:Product {id: 'prod-001'})
			SET p1.name = 'Smartwatch Gen1', p1.sku = 'SKU-0001', p1.price = 1038.25, p1.weight = 1.3, p1.leadTime = 5, p1.status = 'active'

			MERGE (p2:Product {id: 'prod-002'})
			SET p2.name = 'Smartphone Gen2', p2.sku = 'SKU-0002', p2.price = 909.69, p2.weight = 1.1, p2.leadTime = 5, p2.status = 'active'

			MERGE (p3:Product {id: 'prod-003'})
			SET p3.name = 'Server Gen3', p3.sku = 'SKU-0003', p3.price = 1000.97, p3.weight = 3.0, p3.leadTime = 24, p3.status = 'active'

			MERGE (p4:Product {id: 'prod-004'})
			SET p4.name = 'Smartwatch Gen4', p4.sku = 'SKU-0004', p4.price = 994.15, p4.weight = 1.4, p4.leadTime = 6, p4.status = 'active'

			MERGE (p5:Product {id: 'prod-005'})
			SET p5.name = 'Server Gen5', p5.sku = 'SKU-0005', p5.price = 1751.35, p5.weight = 3.4, p5.leadTime = 27, p5.status = 'active'

			MERGE (p6:Product {id: 'prod-006'})
			SET p6.name = 'Server Gen6', p6.sku = 'SKU-0006', p6.price = 1177.06, p6.weight = 4.0, p6.leadTime = 19, p6.status = 'active'

			MERGE (p7:Product {id: 'prod-007'})
			SET p7.name = 'Tablet Gen7', p7.sku = 'SKU-0007', p7.price = 1668.01, p7.weight = 4.7, p7.leadTime = 20, p7.status = 'active'

			MERGE (p8:Product {id: 'prod-008'})
			SET p8.name = 'Tablet Gen8', p8.sku = 'SKU-0008', p8.price = 1107.86, p8.weight = 3.7, p8.leadTime = 12, p8.status = 'active'

			MERGE (p9:Product {id: 'prod-009'})
			SET p9.name = 'Tablet Gen9', p9.sku = 'SKU-0009', p9.price = 895.02, p9.weight = 3.7, p9.leadTime = 14, p9.status = 'active'

			MERGE (p10:Product {id: 'prod-010'})
			SET p10.name = 'Laptop Gen10', p10.sku = 'SKU-0010', p10.price = 273.24, p10.weight = 3.5, p10.leadTime = 16, p10.status = 'active'

			MERGE (p11:Product {id: 'prod-011'})
			SET p11.name = 'Router Gen11', p11.sku = 'SKU-0011', p11.price = 829.98, p11.weight = 2.9, p11.leadTime = 5, p11.status = 'active'

			MERGE (p12:Product {id: 'prod-012'})
			SET p12.name = 'Tablet Gen12', p12.sku = 'SKU-0012', p12.price = 241.68, p12.weight = 1.3, p12.leadTime = 19, p12.status = 'active'

			MERGE (p13:Product {id: 'prod-013'})
			SET p13.name = 'Smartwatch Gen13', p13.sku = 'SKU-0013', p13.price = 1056.49, p13.weight = 2.8, p13.leadTime = 27, p13.status = 'active'

			MERGE (p14:Product {id: 'prod-014'})
			SET p14.name = 'Smartwatch Gen14', p14.sku = 'SKU-0014', p14.price = 486.05, p14.weight = 3.6, p14.leadTime = 20, p14.status = 'active'

			MERGE (p15:Product {id: 'prod-015'})
			SET p15.name = 'Drone Gen15', p15.sku = 'SKU-0015', p15.price = 1866.02, p15.weight = 4.5, p15.leadTime = 12, p15.status = 'active'

			MERGE (c1:Component {id: 'comp-001'})
			SET c1.name = 'CPU Chip V1', c1.price = 6.36, c1.quantity = 0, c1.criticality = 'medium'

			MERGE (c2:Component {id: 'comp-002'})
			SET c2.name = 'Sensor V2', c2.price = 91.46, c2.quantity = 0, c2.criticality = 'medium'

			MERGE (c3:Component {id: 'comp-003'})
			SET c3.name = 'Battery V3', c3.price = 6.97, c3.quantity = 0, c3.criticality = 'medium'

			MERGE (c4:Component {id: 'comp-004'})
			SET c4.name = 'Antenna V4', c4.price = 171.75, c4.quantity = 0, c4.criticality = 'low'

			MERGE (c5:Component {id: 'comp-005'})
			SET c5.name = 'CPU Chip V5', c5.price = 124.88, c5.quantity = 0, c5.criticality = 'low'

			MERGE (c6:Component {id: 'comp-006'})
			SET c6.name = 'RAM Module V6', c6.price = 119.10, c6.quantity = 0, c6.criticality = 'low'

			MERGE (c7:Component {id: 'comp-007'})
			SET c7.name = 'Battery V7', c7.price = 68.35, c7.quantity = 0, c7.criticality = 'medium'

			MERGE (c8:Component {id: 'comp-008'})
			SET c8.name = 'RAM Module V8', c8.price = 68.01, c8.quantity = 0, c8.criticality = 'medium'

			MERGE (c9:Component {id: 'comp-009'})
			SET c9.name = 'Power Supply V9', c9.price = 163.04, c9.quantity = 0, c9.criticality = 'high'

			MERGE (c10:Component {id: 'comp-010'})
			SET c10.name = 'Sensor V10', c10.price = 36.24, c10.quantity = 0, c10.criticality = 'low'

			MERGE (c11:Component {id: 'comp-011'})
			SET c11.name = 'CPU Chip V11', c11.price = 195.05, c11.quantity = 0, c11.criticality = 'medium'

			MERGE (c12:Component {id: 'comp-012'})
			SET c12.name = 'Sensor V12', c12.price = 179.03, c12.quantity = 0, c12.criticality = 'high'

			MERGE (c13:Component {id: 'comp-013'})
			SET c13.name = 'Frame V13', c13.price = 91.44, c13.quantity = 0, c13.criticality = 'low'

			MERGE (c14:Component {id: 'comp-014'})
			SET c14.name = 'Antenna V14', c14.price = 166.42, c14.quantity = 0, c14.criticality = 'high'

			MERGE (c15:Component {id: 'comp-015'})
			SET c15.name = 'Frame V15', c15.price = 121.24, c15.quantity = 0, c15.criticality = 'high'

			MERGE (c16:Component {id: 'comp-016'})
			SET c16.name = 'Power Supply V16', c16.price = 136.92, c16.quantity = 0, c16.criticality = 'low'

			MERGE (c17:Component {id: 'comp-017'})
			SET c17.name = 'Power Supply V17', c17.price = 131.62, c17.quantity = 0, c17.criticality = 'high'

			MERGE (c18:Component {id: 'comp-018'})
			SET c18.name = 'Screen V18', c18.price = 175.04, c18.quantity = 0, c18.criticality = 'high'

			MERGE (c19:Component {id: 'comp-019'})
			SET c19.name = 'Battery V19', c19.price = 168.71, c19.quantity = 0, c19.criticality = 'low'

			MERGE (c20:Component {id: 'comp-020'})
			SET c20.name = 'CPU Chip V20', c20.price = 163.22, c20.quantity = 0, c20.criticality = 'low'

			MERGE (c21:Component {id: 'comp-021'})
			SET c21.name = 'Sensor V21', c21.price = 99.54, c21.quantity = 0, c21.criticality = 'medium'

			MERGE (c22:Component {id: 'comp-022'})
			SET c22.name = 'Screen V22', c22.price = 15.88, c22.quantity = 0, c22.criticality = 'high'

			MERGE (c23:Component {id: 'comp-023'})
			SET c23.name = 'CPU Chip V23', c23.price = 60.58, c23.quantity = 0, c23.criticality = 'low'

			MERGE (c24:Component {id: 'comp-024'})
			SET c24.name = 'Battery V24', c24.price = 87.74, c24.quantity = 0, c24.criticality = 'medium'

			MERGE (c25:Component {id: 'comp-025'})
			SET c25.name = 'Sensor V25', c25.price = 84.30, c25.quantity = 0, c25.criticality = 'high'

			MERGE (c26:Component {id: 'comp-026'})
			SET c26.name = 'Sensor V26', c26.price = 132.44, c26.quantity = 0, c26.criticality = 'low'

			MERGE (c27:Component {id: 'comp-027'})
			SET c27.name = 'Camera V27', c27.price = 157.13, c27.quantity = 0, c27.criticality = 'low'

			MERGE (c28:Component {id: 'comp-028'})
			SET c28.name = 'Antenna V28', c28.price = 10.91, c28.quantity = 0, c28.criticality = 'high'

			MERGE (c29:Component {id: 'comp-029'})
			SET c29.name = 'Screen V29', c29.price = 173.75, c29.quantity = 0, c29.criticality = 'high'

			MERGE (c30:Component {id: 'comp-030'})
			SET c30.name = 'Camera V30', c30.price = 198.13, c30.quantity = 0, c30.criticality = 'high'

			MERGE (c31:Component {id: 'comp-031'})
			SET c31.name = 'Antenna V31', c31.price = 161.03, c31.quantity = 0, c31.criticality = 'low'

			MERGE (c32:Component {id: 'comp-032'})
			SET c32.name = 'Power Supply V32', c32.price = 132.33, c32.quantity = 0, c32.criticality = 'high'

			MERGE (c33:Component {id: 'comp-033'})
			SET c33.name = 'Camera V33', c33.price = 135.71, c33.quantity = 0, c33.criticality = 'medium'

			MERGE (c34:Component {id: 'comp-034'})
			SET c34.name = 'SSD V34', c34.price = 66.82, c34.quantity = 0, c34.criticality = 'low'

			MERGE (c35:Component {id: 'comp-035'})
			SET c35.name = 'SSD V35', c35.price = 92.24, c35.quantity = 0, c35.criticality = 'medium'

			MERGE (c36:Component {id: 'comp-036'})
			SET c36.name = 'CPU Chip V36', c36.price = 152.87, c36.quantity = 0, c36.criticality = 'high'

			MERGE (c37:Component {id: 'comp-037'})
			SET c37.name = 'Sensor V37', c37.price = 91.79, c37.quantity = 0, c37.criticality = 'high'

			MERGE (c38:Component {id: 'comp-038'})
			SET c38.name = 'SSD V38', c38.price = 139.30, c38.quantity = 0, c38.criticality = 'medium'

			MERGE (c39:Component {id: 'comp-039'})
			SET c39.name = 'Battery V39', c39.price = 143.35, c39.quantity = 0, c39.criticality = 'low'

			MERGE (c40:Component {id: 'comp-040'})
			SET c40.name = 'RAM Module V40', c40.price = 192.57, c40.quantity = 0, c40.criticality = 'low'

			MERGE (c41:Component {id: 'comp-041'})
			SET c41.name = 'Camera V41', c41.price = 46.39, c41.quantity = 0, c41.criticality = 'high'

			MERGE (c42:Component {id: 'comp-042'})
			SET c42.name = 'Battery V42', c42.price = 159.41, c42.quantity = 0, c42.criticality = 'high'

			MERGE (c43:Component {id: 'comp-043'})
			SET c43.name = 'RAM Module V43', c43.price = 66.32, c43.quantity = 0, c43.criticality = 'high'

			MERGE (c44:Component {id: 'comp-044'})
			SET c44.name = 'Power Supply V44', c44.price = 194.55, c44.quantity = 0, c44.criticality = 'low'

			MERGE (c45:Component {id: 'comp-045'})
			SET c45.name = 'Battery V45', c45.price = 169.16, c45.quantity = 0, c45.criticality = 'high'

			MERGE (c46:Component {id: 'comp-046'})
			SET c46.name = 'Battery V46', c46.price = 187.04, c46.quantity = 0, c46.criticality = 'high'

			MERGE (c47:Component {id: 'comp-047'})
			SET c47.name = 'Camera V47', c47.price = 133.94, c47.quantity = 0, c47.criticality = 'high'

			MERGE (c48:Component {id: 'comp-048'})
			SET c48.name = 'Screen V48', c48.price = 66.80, c48.quantity = 0, c48.criticality = 'low'

			MERGE (c49:Component {id: 'comp-049'})
			SET c49.name = 'Antenna V49', c49.price = 97.25, c49.quantity = 0, c49.criticality = 'high'

			MERGE (c50:Component {id: 'comp-050'})
			SET c50.name = 'Battery V50', c50.price = 80.20, c50.quantity = 0, c50.criticality = 'low'
		`, nil)
		if err != nil { return nil, err }

		_, err = tx.Run(ctx, `
			MERGE (r1:Route {id: 'route-001'})
			SET r1.name = 'Route 1', r1.distance = 729, r1.estimatedTime = 59, r1.cost = 9945, r1.reliability = 0.95

			MERGE (r2:Route {id: 'route-002'})
			SET r2.name = 'Route 2', r2.distance = 7041, r2.estimatedTime = 27, r2.cost = 392, r2.reliability = 0.87

			MERGE (r3:Route {id: 'route-003'})
			SET r3.name = 'Route 3', r3.distance = 3334, r3.estimatedTime = 69, r3.cost = 6762, r3.reliability = 0.88

			MERGE (r4:Route {id: 'route-004'})
			SET r4.name = 'Route 4', r4.distance = 4915, r4.estimatedTime = 133, r4.cost = 1995, r4.reliability = 0.96

			MERGE (r5:Route {id: 'route-005'})
			SET r5.name = 'Route 5', r5.distance = 7848, r5.estimatedTime = 63, r5.cost = 6698, r5.reliability = 0.82

			MERGE (r6:Route {id: 'route-006'})
			SET r6.name = 'Route 6', r6.distance = 1136, r6.estimatedTime = 161, r6.cost = 9526, r6.reliability = 0.87

			MERGE (r7:Route {id: 'route-007'})
			SET r7.name = 'Route 7', r7.distance = 4589, r7.estimatedTime = 87, r7.cost = 1935, r7.reliability = 0.84

			MERGE (r8:Route {id: 'route-008'})
			SET r8.name = 'Route 8', r8.distance = 2487, r8.estimatedTime = 80, r8.cost = 1373, r8.reliability = 0.95

			MERGE (r9:Route {id: 'route-009'})
			SET r9.name = 'Route 9', r9.distance = 1150, r9.estimatedTime = 17, r9.cost = 1312, r9.reliability = 0.89

			MERGE (r10:Route {id: 'route-010'})
			SET r10.name = 'Route 10', r10.distance = 6982, r10.estimatedTime = 131, r10.cost = 1238, r10.reliability = 0.83

			MERGE (r11:Route {id: 'route-011'})
			SET r11.name = 'Route 11', r11.distance = 2195, r11.estimatedTime = 62, r11.cost = 8059, r11.reliability = 0.94

			MERGE (r12:Route {id: 'route-012'})
			SET r12.name = 'Route 12', r12.distance = 171, r12.estimatedTime = 46, r12.cost = 920, r12.reliability = 0.91

			MERGE (r13:Route {id: 'route-013'})
			SET r13.name = 'Route 13', r13.distance = 273, r13.estimatedTime = 17, r13.cost = 4038, r13.reliability = 0.94

			MERGE (r14:Route {id: 'route-014'})
			SET r14.name = 'Route 14', r14.distance = 3200, r14.estimatedTime = 51, r14.cost = 2867, r14.reliability = 0.85

			MERGE (r15:Route {id: 'route-015'})
			SET r15.name = 'Route 15', r15.distance = 596, r15.estimatedTime = 62, r15.cost = 9812, r15.reliability = 0.89

			MERGE (r16:Route {id: 'route-016'})
			SET r16.name = 'Route 16', r16.distance = 6126, r16.estimatedTime = 34, r16.cost = 2802, r16.reliability = 0.86

			MERGE (r17:Route {id: 'route-017'})
			SET r17.name = 'Route 17', r17.distance = 9152, r17.estimatedTime = 15, r17.cost = 6843, r17.reliability = 0.87

			MERGE (r18:Route {id: 'route-018'})
			SET r18.name = 'Route 18', r18.distance = 9586, r18.estimatedTime = 80, r18.cost = 6638, r18.reliability = 0.92

			MERGE (r19:Route {id: 'route-019'})
			SET r19.name = 'Route 19', r19.distance = 5899, r19.estimatedTime = 164, r19.cost = 7587, r19.reliability = 0.85

			MERGE (r20:Route {id: 'route-020'})
			SET r20.name = 'Route 20', r20.distance = 8505, r20.estimatedTime = 186, r20.cost = 3125, r20.reliability = 0.98

			MERGE (r21:Route {id: 'route-021'})
			SET r21.name = 'Route 21', r21.distance = 8551, r21.estimatedTime = 95, r21.cost = 1934, r21.reliability = 0.90

			MERGE (r22:Route {id: 'route-022'})
			SET r22.name = 'Route 22', r22.distance = 201, r22.estimatedTime = 87, r22.cost = 7830, r22.reliability = 0.84

			MERGE (r23:Route {id: 'route-023'})
			SET r23.name = 'Route 23', r23.distance = 5426, r23.estimatedTime = 174, r23.cost = 5339, r23.reliability = 0.88

			MERGE (r24:Route {id: 'route-024'})
			SET r24.name = 'Route 24', r24.distance = 5076, r24.estimatedTime = 118, r24.cost = 7155, r24.reliability = 0.90

			MERGE (r25:Route {id: 'route-025'})
			SET r25.name = 'Route 25', r25.distance = 7833, r25.estimatedTime = 28, r25.cost = 9382, r25.reliability = 0.87

			MERGE (r26:Route {id: 'route-026'})
			SET r26.name = 'Route 26', r26.distance = 8799, r26.estimatedTime = 193, r26.cost = 1977, r26.reliability = 0.96

			MERGE (r27:Route {id: 'route-027'})
			SET r27.name = 'Route 27', r27.distance = 3160, r27.estimatedTime = 63, r27.cost = 2871, r27.reliability = 0.91

			MERGE (r28:Route {id: 'route-028'})
			SET r28.name = 'Route 28', r28.distance = 5185, r28.estimatedTime = 129, r28.cost = 2120, r28.reliability = 0.88

			MERGE (r29:Route {id: 'route-029'})
			SET r29.name = 'Route 29', r29.distance = 6842, r29.estimatedTime = 104, r29.cost = 7153, r29.reliability = 0.82

			MERGE (r30:Route {id: 'route-030'})
			SET r30.name = 'Route 30', r30.distance = 9470, r30.estimatedTime = 172, r30.cost = 3984, r30.reliability = 0.88

			MERGE (r31:Route {id: 'route-031'})
			SET r31.name = 'Route 31', r31.distance = 9833, r31.estimatedTime = 98, r31.cost = 7267, r31.reliability = 0.94

			MERGE (r32:Route {id: 'route-032'})
			SET r32.name = 'Route 32', r32.distance = 3119, r32.estimatedTime = 100, r32.cost = 8487, r32.reliability = 0.86

			MERGE (r33:Route {id: 'route-033'})
			SET r33.name = 'Route 33', r33.distance = 8038, r33.estimatedTime = 34, r33.cost = 4672, r33.reliability = 0.91

			MERGE (r34:Route {id: 'route-034'})
			SET r34.name = 'Route 34', r34.distance = 4397, r34.estimatedTime = 95, r34.cost = 7219, r34.reliability = 0.90

			MERGE (r35:Route {id: 'route-035'})
			SET r35.name = 'Route 35', r35.distance = 7239, r35.estimatedTime = 163, r35.cost = 9438, r35.reliability = 0.96

			MERGE (r36:Route {id: 'route-036'})
			SET r36.name = 'Route 36', r36.distance = 4289, r36.estimatedTime = 178, r36.cost = 3746, r36.reliability = 0.80

			MERGE (r37:Route {id: 'route-037'})
			SET r37.name = 'Route 37', r37.distance = 476, r37.estimatedTime = 76, r37.cost = 6936, r37.reliability = 0.92

			MERGE (r38:Route {id: 'route-038'})
			SET r38.name = 'Route 38', r38.distance = 7358, r38.estimatedTime = 102, r38.cost = 236, r38.reliability = 0.83

			MERGE (r39:Route {id: 'route-039'})
			SET r39.name = 'Route 39', r39.distance = 9990, r39.estimatedTime = 7, r39.cost = 963, r39.reliability = 0.82

			MERGE (r40:Route {id: 'route-040'})
			SET r40.name = 'Route 40', r40.distance = 1877, r40.estimatedTime = 16, r40.cost = 3361, r40.reliability = 0.98
		`, nil)
		if err != nil { return nil, err }

		_, err = tx.Run(ctx, `
			WITH 1 AS dummy MATCH (p:Product {id: 'prod-001'}), (c:Component {id: 'comp-028'}) MERGE (p)-[:COMPOSED_OF {quantity: 4, position: 1}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-001'}), (c:Component {id: 'comp-035'}) MERGE (p)-[:COMPOSED_OF {quantity: 5, position: 2}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-001'}), (c:Component {id: 'comp-019'}) MERGE (p)-[:COMPOSED_OF {quantity: 2, position: 3}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-001'}), (c:Component {id: 'comp-013'}) MERGE (p)-[:COMPOSED_OF {quantity: 4, position: 4}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-001'}), (c:Component {id: 'comp-046'}) MERGE (p)-[:COMPOSED_OF {quantity: 4, position: 5}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-001'}), (c:Component {id: 'comp-009'}) MERGE (p)-[:COMPOSED_OF {quantity: 1, position: 6}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-001'}), (c:Component {id: 'comp-010'}) MERGE (p)-[:COMPOSED_OF {quantity: 5, position: 7}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-001'}), (c:Component {id: 'comp-032'}) MERGE (p)-[:COMPOSED_OF {quantity: 1, position: 8}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-002'}), (c:Component {id: 'comp-021'}) MERGE (p)-[:COMPOSED_OF {quantity: 5, position: 1}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-002'}), (c:Component {id: 'comp-035'}) MERGE (p)-[:COMPOSED_OF {quantity: 4, position: 2}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-002'}), (c:Component {id: 'comp-032'}) MERGE (p)-[:COMPOSED_OF {quantity: 4, position: 3}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-002'}), (c:Component {id: 'comp-033'}) MERGE (p)-[:COMPOSED_OF {quantity: 3, position: 4}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-002'}), (c:Component {id: 'comp-034'}) MERGE (p)-[:COMPOSED_OF {quantity: 3, position: 5}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-002'}), (c:Component {id: 'comp-029'}) MERGE (p)-[:COMPOSED_OF {quantity: 4, position: 6}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-002'}), (c:Component {id: 'comp-027'}) MERGE (p)-[:COMPOSED_OF {quantity: 2, position: 7}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-003'}), (c:Component {id: 'comp-012'}) MERGE (p)-[:COMPOSED_OF {quantity: 4, position: 1}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-003'}), (c:Component {id: 'comp-018'}) MERGE (p)-[:COMPOSED_OF {quantity: 1, position: 2}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-003'}), (c:Component {id: 'comp-021'}) MERGE (p)-[:COMPOSED_OF {quantity: 3, position: 3}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-003'}), (c:Component {id: 'comp-003'}) MERGE (p)-[:COMPOSED_OF {quantity: 1, position: 4}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-003'}), (c:Component {id: 'comp-009'}) MERGE (p)-[:COMPOSED_OF {quantity: 3, position: 5}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-003'}), (c:Component {id: 'comp-032'}) MERGE (p)-[:COMPOSED_OF {quantity: 5, position: 6}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-003'}), (c:Component {id: 'comp-046'}) MERGE (p)-[:COMPOSED_OF {quantity: 1, position: 7}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-004'}), (c:Component {id: 'comp-022'}) MERGE (p)-[:COMPOSED_OF {quantity: 3, position: 1}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-004'}), (c:Component {id: 'comp-020'}) MERGE (p)-[:COMPOSED_OF {quantity: 5, position: 2}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-004'}), (c:Component {id: 'comp-042'}) MERGE (p)-[:COMPOSED_OF {quantity: 5, position: 3}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-004'}), (c:Component {id: 'comp-030'}) MERGE (p)-[:COMPOSED_OF {quantity: 3, position: 4}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-004'}), (c:Component {id: 'comp-046'}) MERGE (p)-[:COMPOSED_OF {quantity: 5, position: 5}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-004'}), (c:Component {id: 'comp-007'}) MERGE (p)-[:COMPOSED_OF {quantity: 4, position: 6}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-005'}), (c:Component {id: 'comp-009'}) MERGE (p)-[:COMPOSED_OF {quantity: 2, position: 1}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-005'}), (c:Component {id: 'comp-005'}) MERGE (p)-[:COMPOSED_OF {quantity: 4, position: 2}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-005'}), (c:Component {id: 'comp-041'}) MERGE (p)-[:COMPOSED_OF {quantity: 4, position: 3}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-005'}), (c:Component {id: 'comp-032'}) MERGE (p)-[:COMPOSED_OF {quantity: 5, position: 4}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-005'}), (c:Component {id: 'comp-047'}) MERGE (p)-[:COMPOSED_OF {quantity: 4, position: 5}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-006'}), (c:Component {id: 'comp-015'}) MERGE (p)-[:COMPOSED_OF {quantity: 5, position: 1}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-006'}), (c:Component {id: 'comp-009'}) MERGE (p)-[:COMPOSED_OF {quantity: 5, position: 2}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-006'}), (c:Component {id: 'comp-016'}) MERGE (p)-[:COMPOSED_OF {quantity: 3, position: 3}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-006'}), (c:Component {id: 'comp-044'}) MERGE (p)-[:COMPOSED_OF {quantity: 3, position: 4}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-006'}), (c:Component {id: 'comp-011'}) MERGE (p)-[:COMPOSED_OF {quantity: 1, position: 5}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-006'}), (c:Component {id: 'comp-040'}) MERGE (p)-[:COMPOSED_OF {quantity: 3, position: 6}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-006'}), (c:Component {id: 'comp-001'}) MERGE (p)-[:COMPOSED_OF {quantity: 3, position: 7}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-006'}), (c:Component {id: 'comp-041'}) MERGE (p)-[:COMPOSED_OF {quantity: 1, position: 8}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-007'}), (c:Component {id: 'comp-049'}) MERGE (p)-[:COMPOSED_OF {quantity: 5, position: 1}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-007'}), (c:Component {id: 'comp-044'}) MERGE (p)-[:COMPOSED_OF {quantity: 1, position: 2}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-007'}), (c:Component {id: 'comp-037'}) MERGE (p)-[:COMPOSED_OF {quantity: 3, position: 3}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-008'}), (c:Component {id: 'comp-002'}) MERGE (p)-[:COMPOSED_OF {quantity: 1, position: 1}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-008'}), (c:Component {id: 'comp-037'}) MERGE (p)-[:COMPOSED_OF {quantity: 1, position: 2}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-008'}), (c:Component {id: 'comp-021'}) MERGE (p)-[:COMPOSED_OF {quantity: 4, position: 3}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-008'}), (c:Component {id: 'comp-010'}) MERGE (p)-[:COMPOSED_OF {quantity: 1, position: 4}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-008'}), (c:Component {id: 'comp-050'}) MERGE (p)-[:COMPOSED_OF {quantity: 5, position: 5}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-008'}), (c:Component {id: 'comp-025'}) MERGE (p)-[:COMPOSED_OF {quantity: 3, position: 6}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-008'}), (c:Component {id: 'comp-004'}) MERGE (p)-[:COMPOSED_OF {quantity: 2, position: 7}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-009'}), (c:Component {id: 'comp-026'}) MERGE (p)-[:COMPOSED_OF {quantity: 4, position: 1}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-009'}), (c:Component {id: 'comp-048'}) MERGE (p)-[:COMPOSED_OF {quantity: 1, position: 2}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-009'}), (c:Component {id: 'comp-028'}) MERGE (p)-[:COMPOSED_OF {quantity: 1, position: 3}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-009'}), (c:Component {id: 'comp-001'}) MERGE (p)-[:COMPOSED_OF {quantity: 3, position: 4}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-009'}), (c:Component {id: 'comp-029'}) MERGE (p)-[:COMPOSED_OF {quantity: 2, position: 5}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-009'}), (c:Component {id: 'comp-023'}) MERGE (p)-[:COMPOSED_OF {quantity: 5, position: 6}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-010'}), (c:Component {id: 'comp-047'}) MERGE (p)-[:COMPOSED_OF {quantity: 5, position: 1}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-010'}), (c:Component {id: 'comp-020'}) MERGE (p)-[:COMPOSED_OF {quantity: 4, position: 2}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-010'}), (c:Component {id: 'comp-034'}) MERGE (p)-[:COMPOSED_OF {quantity: 5, position: 3}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-010'}), (c:Component {id: 'comp-011'}) MERGE (p)-[:COMPOSED_OF {quantity: 1, position: 4}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-010'}), (c:Component {id: 'comp-024'}) MERGE (p)-[:COMPOSED_OF {quantity: 4, position: 5}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-010'}), (c:Component {id: 'comp-049'}) MERGE (p)-[:COMPOSED_OF {quantity: 3, position: 6}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-010'}), (c:Component {id: 'comp-027'}) MERGE (p)-[:COMPOSED_OF {quantity: 1, position: 7}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-010'}), (c:Component {id: 'comp-040'}) MERGE (p)-[:COMPOSED_OF {quantity: 2, position: 8}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-011'}), (c:Component {id: 'comp-038'}) MERGE (p)-[:COMPOSED_OF {quantity: 4, position: 1}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-011'}), (c:Component {id: 'comp-021'}) MERGE (p)-[:COMPOSED_OF {quantity: 2, position: 2}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-011'}), (c:Component {id: 'comp-041'}) MERGE (p)-[:COMPOSED_OF {quantity: 5, position: 3}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-011'}), (c:Component {id: 'comp-030'}) MERGE (p)-[:COMPOSED_OF {quantity: 4, position: 4}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-012'}), (c:Component {id: 'comp-035'}) MERGE (p)-[:COMPOSED_OF {quantity: 2, position: 1}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-012'}), (c:Component {id: 'comp-023'}) MERGE (p)-[:COMPOSED_OF {quantity: 3, position: 2}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-012'}), (c:Component {id: 'comp-022'}) MERGE (p)-[:COMPOSED_OF {quantity: 3, position: 3}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-012'}), (c:Component {id: 'comp-017'}) MERGE (p)-[:COMPOSED_OF {quantity: 3, position: 4}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-012'}), (c:Component {id: 'comp-032'}) MERGE (p)-[:COMPOSED_OF {quantity: 3, position: 5}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-013'}), (c:Component {id: 'comp-046'}) MERGE (p)-[:COMPOSED_OF {quantity: 5, position: 1}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-013'}), (c:Component {id: 'comp-039'}) MERGE (p)-[:COMPOSED_OF {quantity: 5, position: 2}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-013'}), (c:Component {id: 'comp-033'}) MERGE (p)-[:COMPOSED_OF {quantity: 5, position: 3}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-013'}), (c:Component {id: 'comp-012'}) MERGE (p)-[:COMPOSED_OF {quantity: 2, position: 4}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-013'}), (c:Component {id: 'comp-042'}) MERGE (p)-[:COMPOSED_OF {quantity: 5, position: 5}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-014'}), (c:Component {id: 'comp-034'}) MERGE (p)-[:COMPOSED_OF {quantity: 3, position: 1}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-014'}), (c:Component {id: 'comp-039'}) MERGE (p)-[:COMPOSED_OF {quantity: 3, position: 2}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-014'}), (c:Component {id: 'comp-016'}) MERGE (p)-[:COMPOSED_OF {quantity: 5, position: 3}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-014'}), (c:Component {id: 'comp-047'}) MERGE (p)-[:COMPOSED_OF {quantity: 2, position: 4}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-014'}), (c:Component {id: 'comp-002'}) MERGE (p)-[:COMPOSED_OF {quantity: 1, position: 5}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-015'}), (c:Component {id: 'comp-008'}) MERGE (p)-[:COMPOSED_OF {quantity: 4, position: 1}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-015'}), (c:Component {id: 'comp-034'}) MERGE (p)-[:COMPOSED_OF {quantity: 3, position: 2}]->(c)

			WITH 1 AS dummy MATCH (p:Product {id: 'prod-015'}), (c:Component {id: 'comp-029'}) MERGE (p)-[:COMPOSED_OF {quantity: 3, position: 3}]->(c)
		`, nil)
		if err != nil { return nil, err }

		_, err = tx.Run(ctx, `
			WITH 1 AS dummy MATCH (c:Component {id: 'comp-001'}), (s:Company {id: 'c-025'}) MERGE (c)-[:SUPPLIED_BY {price: 55.70, leadTime: 17, minOrder: 966}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-001'}), (s:Company {id: 'c-021'}) MERGE (c)-[:SUPPLIED_BY {price: 168.03, leadTime: 8, minOrder: 489}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-002'}), (s:Company {id: 'c-021'}) MERGE (c)-[:SUPPLIED_BY {price: 138.03, leadTime: 20, minOrder: 728}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-002'}), (s:Company {id: 'c-025'}) MERGE (c)-[:SUPPLIED_BY {price: 19.54, leadTime: 11, minOrder: 806}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-002'}), (s:Company {id: 'c-012'}) MERGE (c)-[:SUPPLIED_BY {price: 120.24, leadTime: 8, minOrder: 256}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-003'}), (s:Company {id: 'c-020'}) MERGE (c)-[:SUPPLIED_BY {price: 66.42, leadTime: 20, minOrder: 699}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-003'}), (s:Company {id: 'c-002'}) MERGE (c)-[:SUPPLIED_BY {price: 43.31, leadTime: 11, minOrder: 434}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-004'}), (s:Company {id: 'c-007'}) MERGE (c)-[:SUPPLIED_BY {price: 101.96, leadTime: 14, minOrder: 210}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-004'}), (s:Company {id: 'c-020'}) MERGE (c)-[:SUPPLIED_BY {price: 158.01, leadTime: 6, minOrder: 171}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-004'}), (s:Company {id: 'c-005'}) MERGE (c)-[:SUPPLIED_BY {price: 29.01, leadTime: 13, minOrder: 949}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-005'}), (s:Company {id: 'c-016'}) MERGE (c)-[:SUPPLIED_BY {price: 28.33, leadTime: 10, minOrder: 622}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-005'}), (s:Company {id: 'c-025'}) MERGE (c)-[:SUPPLIED_BY {price: 153.49, leadTime: 15, minOrder: 972}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-005'}), (s:Company {id: 'c-005'}) MERGE (c)-[:SUPPLIED_BY {price: 110.29, leadTime: 7, minOrder: 277}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-006'}), (s:Company {id: 'c-025'}) MERGE (c)-[:SUPPLIED_BY {price: 49.86, leadTime: 8, minOrder: 939}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-006'}), (s:Company {id: 'c-016'}) MERGE (c)-[:SUPPLIED_BY {price: 36.11, leadTime: 9, minOrder: 433}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-007'}), (s:Company {id: 'c-021'}) MERGE (c)-[:SUPPLIED_BY {price: 150.51, leadTime: 6, minOrder: 146}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-007'}), (s:Company {id: 'c-012'}) MERGE (c)-[:SUPPLIED_BY {price: 41.28, leadTime: 20, minOrder: 464}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-007'}), (s:Company {id: 'c-020'}) MERGE (c)-[:SUPPLIED_BY {price: 89.14, leadTime: 20, minOrder: 896}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-008'}), (s:Company {id: 'c-024'}) MERGE (c)-[:SUPPLIED_BY {price: 78.46, leadTime: 15, minOrder: 495}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-009'}), (s:Company {id: 'c-011'}) MERGE (c)-[:SUPPLIED_BY {price: 189.82, leadTime: 10, minOrder: 188}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-010'}), (s:Company {id: 'c-022'}) MERGE (c)-[:SUPPLIED_BY {price: 59.07, leadTime: 11, minOrder: 179}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-010'}), (s:Company {id: 'c-016'}) MERGE (c)-[:SUPPLIED_BY {price: 55.68, leadTime: 19, minOrder: 927}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-010'}), (s:Company {id: 'c-018'}) MERGE (c)-[:SUPPLIED_BY {price: 70.05, leadTime: 17, minOrder: 371}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-011'}), (s:Company {id: 'c-021'}) MERGE (c)-[:SUPPLIED_BY {price: 55.34, leadTime: 12, minOrder: 792}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-012'}), (s:Company {id: 'c-011'}) MERGE (c)-[:SUPPLIED_BY {price: 69.37, leadTime: 16, minOrder: 795}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-013'}), (s:Company {id: 'c-030'}) MERGE (c)-[:SUPPLIED_BY {price: 144.20, leadTime: 13, minOrder: 996}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-014'}), (s:Company {id: 'c-021'}) MERGE (c)-[:SUPPLIED_BY {price: 103.31, leadTime: 12, minOrder: 554}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-015'}), (s:Company {id: 'c-009'}) MERGE (c)-[:SUPPLIED_BY {price: 61.39, leadTime: 10, minOrder: 536}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-016'}), (s:Company {id: 'c-025'}) MERGE (c)-[:SUPPLIED_BY {price: 44.82, leadTime: 18, minOrder: 143}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-016'}), (s:Company {id: 'c-022'}) MERGE (c)-[:SUPPLIED_BY {price: 39.74, leadTime: 19, minOrder: 637}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-017'}), (s:Company {id: 'c-012'}) MERGE (c)-[:SUPPLIED_BY {price: 76.04, leadTime: 9, minOrder: 125}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-018'}), (s:Company {id: 'c-016'}) MERGE (c)-[:SUPPLIED_BY {price: 148.96, leadTime: 11, minOrder: 956}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-018'}), (s:Company {id: 'c-022'}) MERGE (c)-[:SUPPLIED_BY {price: 85.03, leadTime: 19, minOrder: 272}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-019'}), (s:Company {id: 'c-025'}) MERGE (c)-[:SUPPLIED_BY {price: 153.89, leadTime: 5, minOrder: 991}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-020'}), (s:Company {id: 'c-021'}) MERGE (c)-[:SUPPLIED_BY {price: 44.14, leadTime: 10, minOrder: 432}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-021'}), (s:Company {id: 'c-030'}) MERGE (c)-[:SUPPLIED_BY {price: 123.11, leadTime: 20, minOrder: 458}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-021'}), (s:Company {id: 'c-018'}) MERGE (c)-[:SUPPLIED_BY {price: 80.25, leadTime: 20, minOrder: 775}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-022'}), (s:Company {id: 'c-011'}) MERGE (c)-[:SUPPLIED_BY {price: 92.54, leadTime: 13, minOrder: 335}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-022'}), (s:Company {id: 'c-007'}) MERGE (c)-[:SUPPLIED_BY {price: 99.18, leadTime: 20, minOrder: 340}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-022'}), (s:Company {id: 'c-020'}) MERGE (c)-[:SUPPLIED_BY {price: 150.75, leadTime: 6, minOrder: 368}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-023'}), (s:Company {id: 'c-024'}) MERGE (c)-[:SUPPLIED_BY {price: 121.73, leadTime: 14, minOrder: 389}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-023'}), (s:Company {id: 'c-016'}) MERGE (c)-[:SUPPLIED_BY {price: 171.83, leadTime: 14, minOrder: 456}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-024'}), (s:Company {id: 'c-012'}) MERGE (c)-[:SUPPLIED_BY {price: 111.04, leadTime: 15, minOrder: 413}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-024'}), (s:Company {id: 'c-025'}) MERGE (c)-[:SUPPLIED_BY {price: 191.44, leadTime: 9, minOrder: 251}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-025'}), (s:Company {id: 'c-022'}) MERGE (c)-[:SUPPLIED_BY {price: 154.48, leadTime: 13, minOrder: 786}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-025'}), (s:Company {id: 'c-009'}) MERGE (c)-[:SUPPLIED_BY {price: 10.01, leadTime: 10, minOrder: 338}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-026'}), (s:Company {id: 'c-005'}) MERGE (c)-[:SUPPLIED_BY {price: 10.17, leadTime: 10, minOrder: 361}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-026'}), (s:Company {id: 'c-021'}) MERGE (c)-[:SUPPLIED_BY {price: 176.38, leadTime: 13, minOrder: 946}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-027'}), (s:Company {id: 'c-022'}) MERGE (c)-[:SUPPLIED_BY {price: 146.50, leadTime: 15, minOrder: 960}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-028'}), (s:Company {id: 'c-021'}) MERGE (c)-[:SUPPLIED_BY {price: 30.46, leadTime: 14, minOrder: 248}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-028'}), (s:Company {id: 'c-007'}) MERGE (c)-[:SUPPLIED_BY {price: 43.90, leadTime: 5, minOrder: 872}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-029'}), (s:Company {id: 'c-025'}) MERGE (c)-[:SUPPLIED_BY {price: 89.52, leadTime: 15, minOrder: 187}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-029'}), (s:Company {id: 'c-009'}) MERGE (c)-[:SUPPLIED_BY {price: 169.19, leadTime: 5, minOrder: 511}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-029'}), (s:Company {id: 'c-011'}) MERGE (c)-[:SUPPLIED_BY {price: 194.12, leadTime: 5, minOrder: 202}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-030'}), (s:Company {id: 'c-025'}) MERGE (c)-[:SUPPLIED_BY {price: 156.82, leadTime: 7, minOrder: 788}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-030'}), (s:Company {id: 'c-021'}) MERGE (c)-[:SUPPLIED_BY {price: 145.14, leadTime: 19, minOrder: 328}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-030'}), (s:Company {id: 'c-012'}) MERGE (c)-[:SUPPLIED_BY {price: 89.92, leadTime: 19, minOrder: 161}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-031'}), (s:Company {id: 'c-011'}) MERGE (c)-[:SUPPLIED_BY {price: 163.61, leadTime: 5, minOrder: 205}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-031'}), (s:Company {id: 'c-016'}) MERGE (c)-[:SUPPLIED_BY {price: 69.40, leadTime: 9, minOrder: 948}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-032'}), (s:Company {id: 'c-007'}) MERGE (c)-[:SUPPLIED_BY {price: 47.62, leadTime: 8, minOrder: 945}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-033'}), (s:Company {id: 'c-021'}) MERGE (c)-[:SUPPLIED_BY {price: 56.38, leadTime: 7, minOrder: 618}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-033'}), (s:Company {id: 'c-009'}) MERGE (c)-[:SUPPLIED_BY {price: 19.87, leadTime: 12, minOrder: 211}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-033'}), (s:Company {id: 'c-018'}) MERGE (c)-[:SUPPLIED_BY {price: 143.87, leadTime: 18, minOrder: 778}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-034'}), (s:Company {id: 'c-012'}) MERGE (c)-[:SUPPLIED_BY {price: 35.70, leadTime: 8, minOrder: 965}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-035'}), (s:Company {id: 'c-007'}) MERGE (c)-[:SUPPLIED_BY {price: 88.64, leadTime: 13, minOrder: 266}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-035'}), (s:Company {id: 'c-025'}) MERGE (c)-[:SUPPLIED_BY {price: 48.20, leadTime: 6, minOrder: 988}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-035'}), (s:Company {id: 'c-005'}) MERGE (c)-[:SUPPLIED_BY {price: 36.83, leadTime: 9, minOrder: 641}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-036'}), (s:Company {id: 'c-025'}) MERGE (c)-[:SUPPLIED_BY {price: 23.10, leadTime: 12, minOrder: 610}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-036'}), (s:Company {id: 'c-002'}) MERGE (c)-[:SUPPLIED_BY {price: 156.70, leadTime: 20, minOrder: 317}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-037'}), (s:Company {id: 'c-009'}) MERGE (c)-[:SUPPLIED_BY {price: 193.48, leadTime: 5, minOrder: 247}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-037'}), (s:Company {id: 'c-012'}) MERGE (c)-[:SUPPLIED_BY {price: 187.29, leadTime: 17, minOrder: 475}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-037'}), (s:Company {id: 'c-020'}) MERGE (c)-[:SUPPLIED_BY {price: 128.57, leadTime: 7, minOrder: 936}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-038'}), (s:Company {id: 'c-005'}) MERGE (c)-[:SUPPLIED_BY {price: 100.44, leadTime: 5, minOrder: 924}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-039'}), (s:Company {id: 'c-025'}) MERGE (c)-[:SUPPLIED_BY {price: 41.33, leadTime: 12, minOrder: 874}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-039'}), (s:Company {id: 'c-020'}) MERGE (c)-[:SUPPLIED_BY {price: 170.10, leadTime: 19, minOrder: 853}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-039'}), (s:Company {id: 'c-030'}) MERGE (c)-[:SUPPLIED_BY {price: 136.80, leadTime: 17, minOrder: 913}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-040'}), (s:Company {id: 'c-025'}) MERGE (c)-[:SUPPLIED_BY {price: 41.26, leadTime: 8, minOrder: 468}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-040'}), (s:Company {id: 'c-005'}) MERGE (c)-[:SUPPLIED_BY {price: 39.19, leadTime: 12, minOrder: 560}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-040'}), (s:Company {id: 'c-018'}) MERGE (c)-[:SUPPLIED_BY {price: 35.15, leadTime: 15, minOrder: 270}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-041'}), (s:Company {id: 'c-030'}) MERGE (c)-[:SUPPLIED_BY {price: 148.15, leadTime: 15, minOrder: 397}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-042'}), (s:Company {id: 'c-002'}) MERGE (c)-[:SUPPLIED_BY {price: 172.66, leadTime: 16, minOrder: 832}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-042'}), (s:Company {id: 'c-009'}) MERGE (c)-[:SUPPLIED_BY {price: 45.67, leadTime: 5, minOrder: 387}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-043'}), (s:Company {id: 'c-009'}) MERGE (c)-[:SUPPLIED_BY {price: 44.00, leadTime: 8, minOrder: 801}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-043'}), (s:Company {id: 'c-012'}) MERGE (c)-[:SUPPLIED_BY {price: 50.48, leadTime: 14, minOrder: 570}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-043'}), (s:Company {id: 'c-016'}) MERGE (c)-[:SUPPLIED_BY {price: 47.07, leadTime: 19, minOrder: 170}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-044'}), (s:Company {id: 'c-024'}) MERGE (c)-[:SUPPLIED_BY {price: 79.87, leadTime: 6, minOrder: 886}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-044'}), (s:Company {id: 'c-005'}) MERGE (c)-[:SUPPLIED_BY {price: 15.67, leadTime: 17, minOrder: 490}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-044'}), (s:Company {id: 'c-002'}) MERGE (c)-[:SUPPLIED_BY {price: 50.92, leadTime: 18, minOrder: 550}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-045'}), (s:Company {id: 'c-020'}) MERGE (c)-[:SUPPLIED_BY {price: 99.83, leadTime: 15, minOrder: 473}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-045'}), (s:Company {id: 'c-018'}) MERGE (c)-[:SUPPLIED_BY {price: 115.88, leadTime: 18, minOrder: 916}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-045'}), (s:Company {id: 'c-009'}) MERGE (c)-[:SUPPLIED_BY {price: 47.29, leadTime: 6, minOrder: 519}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-046'}), (s:Company {id: 'c-018'}) MERGE (c)-[:SUPPLIED_BY {price: 22.91, leadTime: 19, minOrder: 680}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-047'}), (s:Company {id: 'c-002'}) MERGE (c)-[:SUPPLIED_BY {price: 93.48, leadTime: 14, minOrder: 434}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-047'}), (s:Company {id: 'c-025'}) MERGE (c)-[:SUPPLIED_BY {price: 165.61, leadTime: 11, minOrder: 387}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-048'}), (s:Company {id: 'c-020'}) MERGE (c)-[:SUPPLIED_BY {price: 152.24, leadTime: 20, minOrder: 147}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-048'}), (s:Company {id: 'c-025'}) MERGE (c)-[:SUPPLIED_BY {price: 168.13, leadTime: 10, minOrder: 357}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-048'}), (s:Company {id: 'c-005'}) MERGE (c)-[:SUPPLIED_BY {price: 175.25, leadTime: 14, minOrder: 187}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-049'}), (s:Company {id: 'c-002'}) MERGE (c)-[:SUPPLIED_BY {price: 86.50, leadTime: 9, minOrder: 833}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-050'}), (s:Company {id: 'c-016'}) MERGE (c)-[:SUPPLIED_BY {price: 177.67, leadTime: 15, minOrder: 598}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-050'}), (s:Company {id: 'c-009'}) MERGE (c)-[:SUPPLIED_BY {price: 82.01, leadTime: 12, minOrder: 427}]->(s)

			WITH 1 AS dummy MATCH (c:Component {id: 'comp-050'}), (s:Company {id: 'c-022'}) MERGE (c)-[:SUPPLIED_BY {price: 28.41, leadTime: 17, minOrder: 308}]->(s)
		`, nil)
		if err != nil { return nil, err }

		_, err = tx.Run(ctx, `
			MERGE (o1:Order {id: 'order-001'})
			SET o1.orderDate = '2024-01-01', o1.dueDate = '2024-02-01', o1.quantity = 533, o1.status = 'delivered', o1.cost = 30916.95

			WITH 1 AS dummy MATCH (o:Order {id: 'order-001'}), (p:Product {id: 'prod-003'}) MERGE (o)-[:CONTAINS {quantity: 13, unitPrice: 366.87}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-001'}), (c:Company {id: 'c-001'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-001'}), (s:Company {id: 'c-026'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o2:Order {id: 'order-002'})
			SET o2.orderDate = '2024-01-01', o2.dueDate = '2024-02-01', o2.quantity = 915, o2.status = 'pending', o2.cost = 10416.22

			WITH 1 AS dummy MATCH (o:Order {id: 'order-002'}), (p:Product {id: 'prod-015'}) MERGE (o)-[:CONTAINS {quantity: 17, unitPrice: 205.06}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-002'}), (c:Company {id: 'c-013'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-002'}), (s:Company {id: 'c-026'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o3:Order {id: 'order-003'})
			SET o3.orderDate = '2024-01-01', o3.dueDate = '2024-02-01', o3.quantity = 501, o3.status = 'pending', o3.cost = 15394.12

			WITH 1 AS dummy MATCH (o:Order {id: 'order-003'}), (p:Product {id: 'prod-006'}) MERGE (o)-[:CONTAINS {quantity: 65, unitPrice: 75.82}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-003'}), (c:Company {id: 'c-028'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-003'}), (s:Company {id: 'c-010'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o4:Order {id: 'order-004'})
			SET o4.orderDate = '2024-01-01', o4.dueDate = '2024-02-01', o4.quantity = 706, o4.status = 'pending', o4.cost = 66650.69

			WITH 1 AS dummy MATCH (o:Order {id: 'order-004'}), (p:Product {id: 'prod-009'}) MERGE (o)-[:CONTAINS {quantity: 348, unitPrice: 205.93}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-004'}), (c:Company {id: 'c-013'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-004'}), (s:Company {id: 'c-014'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o5:Order {id: 'order-005'})
			SET o5.orderDate = '2024-01-01', o5.dueDate = '2024-02-01', o5.quantity = 877, o5.status = 'in_transit', o5.cost = 91335.49

			WITH 1 AS dummy MATCH (o:Order {id: 'order-005'}), (p:Product {id: 'prod-009'}) MERGE (o)-[:CONTAINS {quantity: 71, unitPrice: 161.62}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-005'}), (c:Company {id: 'c-028'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-005'}), (s:Company {id: 'c-003'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o6:Order {id: 'order-006'})
			SET o6.orderDate = '2024-01-01', o6.dueDate = '2024-02-01', o6.quantity = 182, o6.status = 'pending', o6.cost = 1742.14

			WITH 1 AS dummy MATCH (o:Order {id: 'order-006'}), (p:Product {id: 'prod-004'}) MERGE (o)-[:CONTAINS {quantity: 152, unitPrice: 393.63}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-006'}), (c:Company {id: 'c-004'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-006'}), (s:Company {id: 'c-014'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o7:Order {id: 'order-007'})
			SET o7.orderDate = '2024-01-01', o7.dueDate = '2024-02-01', o7.quantity = 515, o7.status = 'in_transit', o7.cost = 3003.78

			WITH 1 AS dummy MATCH (o:Order {id: 'order-007'}), (p:Product {id: 'prod-010'}) MERGE (o)-[:CONTAINS {quantity: 122, unitPrice: 162.22}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-007'}), (c:Company {id: 'c-029'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-007'}), (s:Company {id: 'c-010'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o8:Order {id: 'order-008'})
			SET o8.orderDate = '2024-01-01', o8.dueDate = '2024-02-01', o8.quantity = 735, o8.status = 'in_transit', o8.cost = 22752.58

			WITH 1 AS dummy MATCH (o:Order {id: 'order-008'}), (p:Product {id: 'prod-001'}) MERGE (o)-[:CONTAINS {quantity: 21, unitPrice: 238.25}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-008'}), (c:Company {id: 'c-029'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-008'}), (s:Company {id: 'c-003'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o9:Order {id: 'order-009'})
			SET o9.orderDate = '2024-01-01', o9.dueDate = '2024-02-01', o9.quantity = 264, o9.status = 'in_transit', o9.cost = 2485.25

			WITH 1 AS dummy MATCH (o:Order {id: 'order-009'}), (p:Product {id: 'prod-006'}) MERGE (o)-[:CONTAINS {quantity: 429, unitPrice: 393.20}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-009'}), (c:Company {id: 'c-029'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-009'}), (s:Company {id: 'c-010'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o10:Order {id: 'order-010'})
			SET o10.orderDate = '2024-01-01', o10.dueDate = '2024-02-01', o10.quantity = 557, o10.status = 'delayed', o10.cost = 56692.36

			WITH 1 AS dummy MATCH (o:Order {id: 'order-010'}), (p:Product {id: 'prod-002'}) MERGE (o)-[:CONTAINS {quantity: 178, unitPrice: 425.94}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-010'}), (c:Company {id: 'c-028'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-010'}), (s:Company {id: 'c-023'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o11:Order {id: 'order-011'})
			SET o11.orderDate = '2024-01-01', o11.dueDate = '2024-02-01', o11.quantity = 975, o11.status = 'delayed', o11.cost = 77710.77

			WITH 1 AS dummy MATCH (o:Order {id: 'order-011'}), (p:Product {id: 'prod-015'}) MERGE (o)-[:CONTAINS {quantity: 329, unitPrice: 101.93}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-011'}), (c:Company {id: 'c-017'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-011'}), (s:Company {id: 'c-023'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o12:Order {id: 'order-012'})
			SET o12.orderDate = '2024-01-01', o12.dueDate = '2024-02-01', o12.quantity = 413, o12.status = 'delivered', o12.cost = 92193.04

			WITH 1 AS dummy MATCH (o:Order {id: 'order-012'}), (p:Product {id: 'prod-012'}) MERGE (o)-[:CONTAINS {quantity: 180, unitPrice: 65.34}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-012'}), (c:Company {id: 'c-006'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-012'}), (s:Company {id: 'c-019'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o13:Order {id: 'order-013'})
			SET o13.orderDate = '2024-01-01', o13.dueDate = '2024-02-01', o13.quantity = 863, o13.status = 'pending', o13.cost = 88602.09

			WITH 1 AS dummy MATCH (o:Order {id: 'order-013'}), (p:Product {id: 'prod-010'}) MERGE (o)-[:CONTAINS {quantity: 447, unitPrice: 342.78}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-013'}), (c:Company {id: 'c-004'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-013'}), (s:Company {id: 'c-003'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o14:Order {id: 'order-014'})
			SET o14.orderDate = '2024-01-01', o14.dueDate = '2024-02-01', o14.quantity = 414, o14.status = 'pending', o14.cost = 52273.66

			WITH 1 AS dummy MATCH (o:Order {id: 'order-014'}), (p:Product {id: 'prod-004'}) MERGE (o)-[:CONTAINS {quantity: 52, unitPrice: 192.21}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-014'}), (c:Company {id: 'c-013'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-014'}), (s:Company {id: 'c-010'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o15:Order {id: 'order-015'})
			SET o15.orderDate = '2024-01-01', o15.dueDate = '2024-02-01', o15.quantity = 747, o15.status = 'pending', o15.cost = 51003.50

			WITH 1 AS dummy MATCH (o:Order {id: 'order-015'}), (p:Product {id: 'prod-007'}) MERGE (o)-[:CONTAINS {quantity: 429, unitPrice: 68.40}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-015'}), (c:Company {id: 'c-001'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-015'}), (s:Company {id: 'c-019'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o16:Order {id: 'order-016'})
			SET o16.orderDate = '2024-01-01', o16.dueDate = '2024-02-01', o16.quantity = 206, o16.status = 'delayed', o16.cost = 7814.62

			WITH 1 AS dummy MATCH (o:Order {id: 'order-016'}), (p:Product {id: 'prod-001'}) MERGE (o)-[:CONTAINS {quantity: 338, unitPrice: 219.83}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-016'}), (c:Company {id: 'c-013'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-016'}), (s:Company {id: 'c-023'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o17:Order {id: 'order-017'})
			SET o17.orderDate = '2024-01-01', o17.dueDate = '2024-02-01', o17.quantity = 530, o17.status = 'delivered', o17.cost = 42839.33

			WITH 1 AS dummy MATCH (o:Order {id: 'order-017'}), (p:Product {id: 'prod-003'}) MERGE (o)-[:CONTAINS {quantity: 89, unitPrice: 474.79}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-017'}), (c:Company {id: 'c-004'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-017'}), (s:Company {id: 'c-023'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o18:Order {id: 'order-018'})
			SET o18.orderDate = '2024-01-01', o18.dueDate = '2024-02-01', o18.quantity = 563, o18.status = 'in_transit', o18.cost = 73077.17

			WITH 1 AS dummy MATCH (o:Order {id: 'order-018'}), (p:Product {id: 'prod-009'}) MERGE (o)-[:CONTAINS {quantity: 390, unitPrice: 302.61}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-018'}), (c:Company {id: 'c-029'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-018'}), (s:Company {id: 'c-003'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o19:Order {id: 'order-019'})
			SET o19.orderDate = '2024-01-01', o19.dueDate = '2024-02-01', o19.quantity = 593, o19.status = 'delayed', o19.cost = 42878.07

			WITH 1 AS dummy MATCH (o:Order {id: 'order-019'}), (p:Product {id: 'prod-015'}) MERGE (o)-[:CONTAINS {quantity: 374, unitPrice: 97.85}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-019'}), (c:Company {id: 'c-028'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-019'}), (s:Company {id: 'c-003'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o20:Order {id: 'order-020'})
			SET o20.orderDate = '2024-01-01', o20.dueDate = '2024-02-01', o20.quantity = 299, o20.status = 'delayed', o20.cost = 15117.61

			WITH 1 AS dummy MATCH (o:Order {id: 'order-020'}), (p:Product {id: 'prod-003'}) MERGE (o)-[:CONTAINS {quantity: 241, unitPrice: 286.26}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-020'}), (c:Company {id: 'c-004'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-020'}), (s:Company {id: 'c-019'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o21:Order {id: 'order-021'})
			SET o21.orderDate = '2024-01-01', o21.dueDate = '2024-02-01', o21.quantity = 608, o21.status = 'delayed', o21.cost = 47117.45

			WITH 1 AS dummy MATCH (o:Order {id: 'order-021'}), (p:Product {id: 'prod-010'}) MERGE (o)-[:CONTAINS {quantity: 242, unitPrice: 95.59}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-021'}), (c:Company {id: 'c-029'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-021'}), (s:Company {id: 'c-014'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o22:Order {id: 'order-022'})
			SET o22.orderDate = '2024-01-01', o22.dueDate = '2024-02-01', o22.quantity = 262, o22.status = 'delayed', o22.cost = 4513.06

			WITH 1 AS dummy MATCH (o:Order {id: 'order-022'}), (p:Product {id: 'prod-005'}) MERGE (o)-[:CONTAINS {quantity: 129, unitPrice: 195.64}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-022'}), (c:Company {id: 'c-017'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-022'}), (s:Company {id: 'c-027'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o23:Order {id: 'order-023'})
			SET o23.orderDate = '2024-01-01', o23.dueDate = '2024-02-01', o23.quantity = 588, o23.status = 'delivered', o23.cost = 15544.60

			WITH 1 AS dummy MATCH (o:Order {id: 'order-023'}), (p:Product {id: 'prod-011'}) MERGE (o)-[:CONTAINS {quantity: 148, unitPrice: 439.21}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-023'}), (c:Company {id: 'c-015'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-023'}), (s:Company {id: 'c-014'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o24:Order {id: 'order-024'})
			SET o24.orderDate = '2024-01-01', o24.dueDate = '2024-02-01', o24.quantity = 802, o24.status = 'delayed', o24.cost = 58758.85

			WITH 1 AS dummy MATCH (o:Order {id: 'order-024'}), (p:Product {id: 'prod-015'}) MERGE (o)-[:CONTAINS {quantity: 251, unitPrice: 409.05}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-024'}), (c:Company {id: 'c-029'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-024'}), (s:Company {id: 'c-023'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o25:Order {id: 'order-025'})
			SET o25.orderDate = '2024-01-01', o25.dueDate = '2024-02-01', o25.quantity = 798, o25.status = 'delivered', o25.cost = 67855.31

			WITH 1 AS dummy MATCH (o:Order {id: 'order-025'}), (p:Product {id: 'prod-014'}) MERGE (o)-[:CONTAINS {quantity: 35, unitPrice: 203.05}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-025'}), (c:Company {id: 'c-004'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-025'}), (s:Company {id: 'c-019'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o26:Order {id: 'order-026'})
			SET o26.orderDate = '2024-01-01', o26.dueDate = '2024-02-01', o26.quantity = 925, o26.status = 'delayed', o26.cost = 76829.54

			WITH 1 AS dummy MATCH (o:Order {id: 'order-026'}), (p:Product {id: 'prod-001'}) MERGE (o)-[:CONTAINS {quantity: 114, unitPrice: 495.09}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-026'}), (c:Company {id: 'c-028'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-026'}), (s:Company {id: 'c-014'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o27:Order {id: 'order-027'})
			SET o27.orderDate = '2024-01-01', o27.dueDate = '2024-02-01', o27.quantity = 119, o27.status = 'delivered', o27.cost = 6258.81

			WITH 1 AS dummy MATCH (o:Order {id: 'order-027'}), (p:Product {id: 'prod-005'}) MERGE (o)-[:CONTAINS {quantity: 72, unitPrice: 331.88}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-027'}), (c:Company {id: 'c-015'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-027'}), (s:Company {id: 'c-019'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o28:Order {id: 'order-028'})
			SET o28.orderDate = '2024-01-01', o28.dueDate = '2024-02-01', o28.quantity = 794, o28.status = 'pending', o28.cost = 25221.65

			WITH 1 AS dummy MATCH (o:Order {id: 'order-028'}), (p:Product {id: 'prod-011'}) MERGE (o)-[:CONTAINS {quantity: 262, unitPrice: 245.46}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-028'}), (c:Company {id: 'c-006'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-028'}), (s:Company {id: 'c-027'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o29:Order {id: 'order-029'})
			SET o29.orderDate = '2024-01-01', o29.dueDate = '2024-02-01', o29.quantity = 454, o29.status = 'delivered', o29.cost = 28111.19

			WITH 1 AS dummy MATCH (o:Order {id: 'order-029'}), (p:Product {id: 'prod-001'}) MERGE (o)-[:CONTAINS {quantity: 349, unitPrice: 206.79}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-029'}), (c:Company {id: 'c-001'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-029'}), (s:Company {id: 'c-014'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o30:Order {id: 'order-030'})
			SET o30.orderDate = '2024-01-01', o30.dueDate = '2024-02-01', o30.quantity = 947, o30.status = 'pending', o30.cost = 24268.60

			WITH 1 AS dummy MATCH (o:Order {id: 'order-030'}), (p:Product {id: 'prod-012'}) MERGE (o)-[:CONTAINS {quantity: 498, unitPrice: 396.49}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-030'}), (c:Company {id: 'c-028'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-030'}), (s:Company {id: 'c-023'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o31:Order {id: 'order-031'})
			SET o31.orderDate = '2024-01-01', o31.dueDate = '2024-02-01', o31.quantity = 194, o31.status = 'in_transit', o31.cost = 53243.21

			WITH 1 AS dummy MATCH (o:Order {id: 'order-031'}), (p:Product {id: 'prod-008'}) MERGE (o)-[:CONTAINS {quantity: 211, unitPrice: 142.47}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-031'}), (c:Company {id: 'c-029'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-031'}), (s:Company {id: 'c-027'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o32:Order {id: 'order-032'})
			SET o32.orderDate = '2024-01-01', o32.dueDate = '2024-02-01', o32.quantity = 639, o32.status = 'delivered', o32.cost = 78719.00

			WITH 1 AS dummy MATCH (o:Order {id: 'order-032'}), (p:Product {id: 'prod-004'}) MERGE (o)-[:CONTAINS {quantity: 232, unitPrice: 199.63}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-032'}), (c:Company {id: 'c-015'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-032'}), (s:Company {id: 'c-003'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o33:Order {id: 'order-033'})
			SET o33.orderDate = '2024-01-01', o33.dueDate = '2024-02-01', o33.quantity = 834, o33.status = 'delivered', o33.cost = 6985.80

			WITH 1 AS dummy MATCH (o:Order {id: 'order-033'}), (p:Product {id: 'prod-003'}) MERGE (o)-[:CONTAINS {quantity: 75, unitPrice: 155.87}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-033'}), (c:Company {id: 'c-017'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-033'}), (s:Company {id: 'c-019'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o34:Order {id: 'order-034'})
			SET o34.orderDate = '2024-01-01', o34.dueDate = '2024-02-01', o34.quantity = 665, o34.status = 'in_transit', o34.cost = 69714.67

			WITH 1 AS dummy MATCH (o:Order {id: 'order-034'}), (p:Product {id: 'prod-006'}) MERGE (o)-[:CONTAINS {quantity: 189, unitPrice: 308.30}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-034'}), (c:Company {id: 'c-006'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-034'}), (s:Company {id: 'c-027'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o35:Order {id: 'order-035'})
			SET o35.orderDate = '2024-01-01', o35.dueDate = '2024-02-01', o35.quantity = 196, o35.status = 'pending', o35.cost = 23395.79

			WITH 1 AS dummy MATCH (o:Order {id: 'order-035'}), (p:Product {id: 'prod-015'}) MERGE (o)-[:CONTAINS {quantity: 321, unitPrice: 452.14}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-035'}), (c:Company {id: 'c-029'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-035'}), (s:Company {id: 'c-010'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o36:Order {id: 'order-036'})
			SET o36.orderDate = '2024-01-01', o36.dueDate = '2024-02-01', o36.quantity = 673, o36.status = 'delayed', o36.cost = 66858.39

			WITH 1 AS dummy MATCH (o:Order {id: 'order-036'}), (p:Product {id: 'prod-009'}) MERGE (o)-[:CONTAINS {quantity: 386, unitPrice: 63.11}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-036'}), (c:Company {id: 'c-029'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-036'}), (s:Company {id: 'c-014'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o37:Order {id: 'order-037'})
			SET o37.orderDate = '2024-01-01', o37.dueDate = '2024-02-01', o37.quantity = 554, o37.status = 'pending', o37.cost = 19416.47

			WITH 1 AS dummy MATCH (o:Order {id: 'order-037'}), (p:Product {id: 'prod-015'}) MERGE (o)-[:CONTAINS {quantity: 458, unitPrice: 311.38}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-037'}), (c:Company {id: 'c-008'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-037'}), (s:Company {id: 'c-010'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o38:Order {id: 'order-038'})
			SET o38.orderDate = '2024-01-01', o38.dueDate = '2024-02-01', o38.quantity = 700, o38.status = 'delayed', o38.cost = 70546.83

			WITH 1 AS dummy MATCH (o:Order {id: 'order-038'}), (p:Product {id: 'prod-012'}) MERGE (o)-[:CONTAINS {quantity: 105, unitPrice: 134.64}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-038'}), (c:Company {id: 'c-028'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-038'}), (s:Company {id: 'c-019'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o39:Order {id: 'order-039'})
			SET o39.orderDate = '2024-01-01', o39.dueDate = '2024-02-01', o39.quantity = 851, o39.status = 'in_transit', o39.cost = 21063.45

			WITH 1 AS dummy MATCH (o:Order {id: 'order-039'}), (p:Product {id: 'prod-002'}) MERGE (o)-[:CONTAINS {quantity: 129, unitPrice: 140.15}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-039'}), (c:Company {id: 'c-004'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-039'}), (s:Company {id: 'c-019'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o40:Order {id: 'order-040'})
			SET o40.orderDate = '2024-01-01', o40.dueDate = '2024-02-01', o40.quantity = 241, o40.status = 'delayed', o40.cost = 43860.16

			WITH 1 AS dummy MATCH (o:Order {id: 'order-040'}), (p:Product {id: 'prod-001'}) MERGE (o)-[:CONTAINS {quantity: 250, unitPrice: 320.83}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-040'}), (c:Company {id: 'c-004'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-040'}), (s:Company {id: 'c-019'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o41:Order {id: 'order-041'})
			SET o41.orderDate = '2024-01-01', o41.dueDate = '2024-02-01', o41.quantity = 290, o41.status = 'delayed', o41.cost = 35495.52

			WITH 1 AS dummy MATCH (o:Order {id: 'order-041'}), (p:Product {id: 'prod-013'}) MERGE (o)-[:CONTAINS {quantity: 107, unitPrice: 332.92}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-041'}), (c:Company {id: 'c-015'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-041'}), (s:Company {id: 'c-003'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o42:Order {id: 'order-042'})
			SET o42.orderDate = '2024-01-01', o42.dueDate = '2024-02-01', o42.quantity = 95, o42.status = 'pending', o42.cost = 91255.61

			WITH 1 AS dummy MATCH (o:Order {id: 'order-042'}), (p:Product {id: 'prod-010'}) MERGE (o)-[:CONTAINS {quantity: 499, unitPrice: 71.94}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-042'}), (c:Company {id: 'c-006'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-042'}), (s:Company {id: 'c-003'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o43:Order {id: 'order-043'})
			SET o43.orderDate = '2024-01-01', o43.dueDate = '2024-02-01', o43.quantity = 90, o43.status = 'delayed', o43.cost = 81160.45

			WITH 1 AS dummy MATCH (o:Order {id: 'order-043'}), (p:Product {id: 'prod-010'}) MERGE (o)-[:CONTAINS {quantity: 132, unitPrice: 251.42}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-043'}), (c:Company {id: 'c-017'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-043'}), (s:Company {id: 'c-026'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o44:Order {id: 'order-044'})
			SET o44.orderDate = '2024-01-01', o44.dueDate = '2024-02-01', o44.quantity = 802, o44.status = 'in_transit', o44.cost = 88744.10

			WITH 1 AS dummy MATCH (o:Order {id: 'order-044'}), (p:Product {id: 'prod-010'}) MERGE (o)-[:CONTAINS {quantity: 80, unitPrice: 217.72}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-044'}), (c:Company {id: 'c-008'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-044'}), (s:Company {id: 'c-023'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o45:Order {id: 'order-045'})
			SET o45.orderDate = '2024-01-01', o45.dueDate = '2024-02-01', o45.quantity = 218, o45.status = 'delivered', o45.cost = 55036.34

			WITH 1 AS dummy MATCH (o:Order {id: 'order-045'}), (p:Product {id: 'prod-006'}) MERGE (o)-[:CONTAINS {quantity: 455, unitPrice: 285.98}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-045'}), (c:Company {id: 'c-004'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-045'}), (s:Company {id: 'c-003'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o46:Order {id: 'order-046'})
			SET o46.orderDate = '2024-01-01', o46.dueDate = '2024-02-01', o46.quantity = 507, o46.status = 'delayed', o46.cost = 23302.66

			WITH 1 AS dummy MATCH (o:Order {id: 'order-046'}), (p:Product {id: 'prod-012'}) MERGE (o)-[:CONTAINS {quantity: 459, unitPrice: 190.80}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-046'}), (c:Company {id: 'c-015'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-046'}), (s:Company {id: 'c-010'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o47:Order {id: 'order-047'})
			SET o47.orderDate = '2024-01-01', o47.dueDate = '2024-02-01', o47.quantity = 337, o47.status = 'pending', o47.cost = 90676.87

			WITH 1 AS dummy MATCH (o:Order {id: 'order-047'}), (p:Product {id: 'prod-010'}) MERGE (o)-[:CONTAINS {quantity: 140, unitPrice: 254.42}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-047'}), (c:Company {id: 'c-028'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-047'}), (s:Company {id: 'c-014'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o48:Order {id: 'order-048'})
			SET o48.orderDate = '2024-01-01', o48.dueDate = '2024-02-01', o48.quantity = 875, o48.status = 'delayed', o48.cost = 12445.46

			WITH 1 AS dummy MATCH (o:Order {id: 'order-048'}), (p:Product {id: 'prod-009'}) MERGE (o)-[:CONTAINS {quantity: 499, unitPrice: 127.86}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-048'}), (c:Company {id: 'c-029'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-048'}), (s:Company {id: 'c-026'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o49:Order {id: 'order-049'})
			SET o49.orderDate = '2024-01-01', o49.dueDate = '2024-02-01', o49.quantity = 54, o49.status = 'delayed', o49.cost = 22402.13

			WITH 1 AS dummy MATCH (o:Order {id: 'order-049'}), (p:Product {id: 'prod-011'}) MERGE (o)-[:CONTAINS {quantity: 315, unitPrice: 304.02}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-049'}), (c:Company {id: 'c-008'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-049'}), (s:Company {id: 'c-003'}) MERGE (o)-[:PLACED_WITH]->(s)

			MERGE (o50:Order {id: 'order-050'})
			SET o50.orderDate = '2024-01-01', o50.dueDate = '2024-02-01', o50.quantity = 981, o50.status = 'pending', o50.cost = 25025.41

			WITH 1 AS dummy MATCH (o:Order {id: 'order-050'}), (p:Product {id: 'prod-009'}) MERGE (o)-[:CONTAINS {quantity: 128, unitPrice: 280.36}]->(p)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-050'}), (c:Company {id: 'c-028'}) MERGE (o)-[:FROM]->(c)

			WITH 1 AS dummy MATCH (o:Order {id: 'order-050'}), (s:Company {id: 'c-027'}) MERGE (o)-[:PLACED_WITH]->(s)
		`, nil)
		if err != nil { return nil, err }


		return nil, nil
	})

	if err != nil {
		zap.L().Error("Error seeding database", zap.Error(err))
		return err
	}
	zap.L().Info("Database seeded successfully")
	return nil
}

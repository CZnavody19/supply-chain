import random

NUM_COMPANIES = 30
NUM_PRODUCTS = 15
NUM_COMPONENTS = 50
NUM_LOCATIONS = 20
NUM_ROUTES = 40
NUM_ORDERS = 50

countries = [("Taiwan", 25.0, 121.5), ("Germany", 51.1, 10.4), ("China", 35.8, 104.1), ("USA", 37.0, -95.7), ("Czech Republic", 49.8, 15.4), ("Japan", 36.2, 138.2), ("Vietnam", 14.0, 108.2)]
comp_types = ["supplier", "supplier", "supplier", "manufacturer", "manufacturer", "distributor", "retailer", "customer"]
prod_names = ["Laptop", "Smartphone", "Tablet", "Server", "Smartwatch", "Drone", "Monitor", "Router"]
comp_names = ["CPU Chip", "RAM Module", "SSD", "Battery", "Screen", "Frame", "Sensor", "Camera", "Antenna", "Power Supply"]
loc_types = ["warehouse", "warehouse", "distribution_center", "port"]

def generate_cypher_blocks():
    go_code = """package db

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
	"go.uber.org/zap"
)

func (ds *DatabaseStore) SeedDatabase(ctx context.Context) error {
	session := ds.newSession(ctx)
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
"""
    def add_block(queries):
        nonlocal go_code
        go_code += "\t\t_, err = tx.Run(ctx, `\n" + "\n\n".join(queries) + "\n\t\t`, nil)\n\t\tif err != nil { return nil, err }\n\n"

    # Companies
    queries = []
    companies = []
    for i in range(1, NUM_COMPANIES + 1):
        c_id = f"c-{i:03d}"
        ctype = random.choice(comp_types)
        country, lat, lng = random.choice(countries)
        lat += random.uniform(-1, 1)
        lng += random.uniform(-1, 1)
        companies.append((c_id, ctype))
        queries.append(f"\t\t\tMERGE (c{i}:Company {{id: '{c_id}'}})\n\t\t\tSET c{i}.name = 'Company {i}', c{i}.type = '{ctype}', c{i}.country = '{country}', c{i}.lat = {lat:.3f}, c{i}.lng = {lng:.3f}, c{i}.reliability = {random.uniform(0.7, 1.0):.2f}")
    add_block(queries)

    # Locations
    queries = []
    for i in range(1, NUM_LOCATIONS + 1):
        l_id = f"loc-{i:03d}"
        ltype = random.choice(loc_types)
        _, lat, lng = random.choice(countries)
        queries.append(f"\t\t\tMERGE (l{i}:Location {{id: '{l_id}'}})\n\t\t\tSET l{i}.name = 'Location {i}', l{i}.type = '{ltype}', l{i}.lat = {lat:.3f}, l{i}.lng = {lng:.3f}, l{i}.capacity = {random.randint(1000, 50000)}")
    add_block(queries)

    # Products & Components
    queries = []
    for i in range(1, NUM_PRODUCTS + 1):
        queries.append(f"\t\t\tMERGE (p{i}:Product {{id: 'prod-{i:03d}'}})\n\t\t\tSET p{i}.name = '{random.choice(prod_names)} Gen{i}', p{i}.sku = 'SKU-{i:04d}', p{i}.price = {random.uniform(100, 2000):.2f}, p{i}.weight = {random.uniform(0.1, 5.0):.1f}, p{i}.leadTime = {random.randint(5, 30)}, p{i}.status = 'active'")
    for i in range(1, NUM_COMPONENTS + 1):
        queries.append(f"\t\t\tMERGE (c{i}:Component {{id: 'comp-{i:03d}'}})\n\t\t\tSET c{i}.name = '{random.choice(comp_names)} V{i}', c{i}.price = {random.uniform(5, 200):.2f}, c{i}.quantity = 0, c{i}.criticality = '{random.choice(['low', 'medium', 'high'])}'")
    add_block(queries)

    # Routes
    queries = []
    for i in range(1, NUM_ROUTES + 1):
        queries.append(f"\t\t\tMERGE (r{i}:Route {{id: 'route-{i:03d}'}})\n\t\t\tSET r{i}.name = 'Route {i}', r{i}.distance = {random.randint(100, 10000)}, r{i}.estimatedTime = {random.randint(5, 200)}, r{i}.cost = {random.randint(200, 10000)}, r{i}.reliability = {random.uniform(0.8, 0.99):.2f}")
    add_block(queries)

    # BOM (COMPOSED_OF)
    queries = []
    for p in range(1, NUM_PRODUCTS + 1):
        comps = random.sample(range(1, NUM_COMPONENTS + 1), random.randint(3, 8))
        for pos, c in enumerate(comps, 1):
            queries.append(f"\t\t\tWITH 1 AS dummy MATCH (p:Product {{id: 'prod-{p:03d}'}}), (c:Component {{id: 'comp-{c:03d}'}}) MERGE (p)-[:COMPOSED_OF {{quantity: {random.randint(1, 5)}, position: {pos}}}]->(c)")
    add_block(queries)

    # SUPPLIED_BY (Components -> Suppliers)
    queries = []
    suppliers = [c[0] for c in companies if c[1] == 'supplier']
    for i in range(1, NUM_COMPONENTS + 1):
        sups = random.sample(suppliers, random.randint(1, 3))
        for s in sups:
            queries.append(f"\t\t\tWITH 1 AS dummy MATCH (c:Component {{id: 'comp-{i:03d}'}}), (s:Company {{id: '{s}'}}) MERGE (c)-[:SUPPLIED_BY {{price: {random.uniform(10, 200):.2f}, leadTime: {random.randint(5, 20)}, minOrder: {random.randint(100, 1000)}}}]->(s)")
    add_block(queries)

    # Orders
    queries = []
    for i in range(1, NUM_ORDERS + 1):
        status = random.choice(['pending', 'in_transit', 'delivered', 'delayed'])
        queries.append(f"\t\t\tMERGE (o{i}:Order {{id: 'order-{i:03d}'}})\n\t\t\tSET o{i}.orderDate = '2024-01-01', o{i}.dueDate = '2024-02-01', o{i}.quantity = {random.randint(10, 1000)}, o{i}.status = '{status}', o{i}.cost = {random.uniform(1000, 100000):.2f}")
        
        prod_id = f"prod-{random.randint(1, NUM_PRODUCTS):03d}"
        cust = random.choice([c[0] for c in companies if c[1] in ('customer', 'retailer')])
        sup = random.choice([c[0] for c in companies if c[1] in ('manufacturer', 'distributor')])
        queries.append(f"\t\t\tWITH 1 AS dummy MATCH (o:Order {{id: 'order-{i:03d}'}}), (p:Product {{id: '{prod_id}'}}) MERGE (o)-[:CONTAINS {{quantity: {random.randint(10, 500)}, unitPrice: {random.uniform(50, 500):.2f}}}]->(p)")
        queries.append(f"\t\t\tWITH 1 AS dummy MATCH (o:Order {{id: 'order-{i:03d}'}}), (c:Company {{id: '{cust}'}}) MERGE (o)-[:FROM]->(c)")
        queries.append(f"\t\t\tWITH 1 AS dummy MATCH (o:Order {{id: 'order-{i:03d}'}}), (s:Company {{id: '{sup}'}}) MERGE (o)-[:PLACED_WITH]->(s)")
    add_block(queries)

    go_code += """
		return nil, nil
	})

	if err != nil {
		zap.L().Error("Error seeding database", zap.Error(err))
		return err
	}
	zap.L().Info("Database seeded successfully")
	return nil
}
"""
    go_code = go_code.replace("\t\t_, err = tx.Run(ctx, `", "\t\tvar err error\n\t\t_, err = tx.Run(ctx, `", 1)
    return go_code

if __name__ == "__main__":
    with open("seed.go", "w") as f:
        f.write(generate_cypher_blocks())
    print("seed.go generated successfully.")
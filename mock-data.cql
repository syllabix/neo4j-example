
CREATE (c119:Condition {id: 7312700156868474266, bolref: 'VMU22049', value: 2.28, type: 'percent', ean: 347732501338814, start: '2018-11-26', end: '2019-03-07'})
// Test Data for CDN DB
CREATE (Samsung:Supplier {name:'Samsung', id:123124 })
CREATE (ca:Condition {id: 12342, name: 'cdn.12342', bolref: 'VET1238912', value: 3.2, type: 'percent', start: '2018-11-01', end: '2018-11-31'})
CREATE (cb:Condition {id: 47893, name: 'cbn.47893', bolref: 'VEE2349823', value: 0.13, type: 'fixed', start: '2018-11-15', end: '2018-12-10'})
CREATE (cc:Condition {id: 38234, name: 'cbn.38234', bolref: 'VTT2349823', value: 0.9, type: 'fixed', start: '2018-12-15', end: '2018-12-31'})

CREATE (sa:ShopOrder {id: 342543, value: 24.34, currency: 'EUR', ean: 'EAN123129123'}),
    (sb:ShopOrder {id: 849345, value: 362.34, currency: 'EUR', ean: 'EAN1200029123,'}),
    (sc:ShopOrder {id: 123345, value: 123.24, currency: 'EUR', ean: 'EAN9230234'}),
    (sd:ShopOrder {id: 342533, value: 98.76, currency: 'EUR', ean: 'EAN9999999'}),
    (se:ShopOrder {id: 234123, value: 698.84, currency: 'EUR', ean: 'EAN82394234'}),
    (sf:ShopOrder {id: 789234, value: 234.34, currency: 'EUR', ean: 'EAN2349238'}),
    (sg:ShopOrder {id: 222334, value: 34.34, currency: 'EUR', ean: 'EAN2388883333'}),
    (sh:ShopOrder {id: 882231, value: 65.67, currency: 'EUR', ean: 'EAN2388883333'})

CREATE
    (Samsung)-[:AGREED_TO]->(ca),
    (Samsung)-[:AGREED_TO]->(cb),
    (Samsung)-[:AGREED_TO]->(cc),
    (sa)-[:MATCHED_TO]->(ca),
    (sb)-[:MATCHED_TO]->(ca),
    (sc)-[:MATCHED_TO]->(ca),
    (sd)-[:MATCHED_TO]->(cb),
    (se)-[:MATCHED_TO]->(cb),
    (sf)-[:MATCHED_TO]->(cc),
    (sg)-[:MATCHED_TO]->(cc),
    (sh)-[:MATCHED_TO]->(cc)

CREATE (Apple:Supplier {name:'Apple', id:12434})
CREATE (caa:Condition {id: 12323, name: 'cdn.12323', bolref: 'VET1238980', value: 6.2, type: 'percent', start: '2018-11-01', end: '2018-11-31'})
CREATE (cab:Condition {id: 94303, name: 'cdn.94303', bolref: 'VEE2342342', value: 0.20, type: 'fixed', start: '2018-11-15', end: '2018-12-10'})
CREATE (cac:Condition {id: 23782, name: 'cdn.23782', bolref: 'VTT2348394', value: 0.11, type: 'fixed', start: '2018-12-15', end: '2018-12-31'})

CREATE (saa:ShopOrder {id: 342543, value: 24.34, currency: 'EUR', ean: 'EAN123129123'})
CREATE (sab:ShopOrder {id: 849345, value: 362.34, currency: 'EUR', ean: 'EAN1200029123'})
CREATE (sac:ShopOrder {id: 123345, value: 123.24, currency: 'EUR', ean: 'EAN9230234'})
CREATE (sad:ShopOrder {id: 342533, value: 98.76, currency: 'EUR', ean: 'EAN9999999'})
CREATE (sae:ShopOrder {id: 234123, value: 698.84, currency: 'EUR', ean: 'EAN82394234'})
CREATE (saf:ShopOrder {id: 789234, value: 234.34, currency: 'EUR', ean: 'EAN2349238'})
CREATE (sag:ShopOrder {id: 667334, value: 34.34, currency: 'EUR', ean: 'EAN2388883333'})
CREATE (sah:ShopOrder {id: 882231, value: 65.67, currency: 'EUR', ean: 'EAN2388883333'})

CREATE
    (Apple)-[:AGREED_TO]->(caa),
    (Apple)-[:AGREED_TO]->(cab),
    (Apple)-[:AGREED_TO]->(cac),
    (saa)-[:MATCHED_TO]->(caa),
    (sab)-[:MATCHED_TO]->(caa),
    (sac)-[:MATCHED_TO]->(caa),
    (sad)-[:MATCHED_TO]->(cab),
    (sae)-[:MATCHED_TO]->(cab),
    (saf)-[:MATCHED_TO]->(cac),
    (sag)-[:MATCHED_TO]->(cac),
    (sah)-[:MATCHED_TO]->(cac)
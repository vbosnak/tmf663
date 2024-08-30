
CREATE TABLE ShoppingCart
(
  id                     INTEGER     NOT NULL,
  href                   TEXT       ,
  creation_date          TIMESTAMP  ,
  lastupdate             TIMESTAMP  ,
  validFor_startDateTime DATE       ,
  validFor_endDateTime   DATE       ,
  cart_price             FLOAT,
  cart_price_id          INTEGER NOT NULL,
  contact_medium_id      INTEGER NOT NULL,
  product_relatedParty_id        INTEGER     NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE CartItem
(
  id                     INTEGER     NOT NULL,
  action                 TEXT       ,
  status                 TEXT       ,
  quantity               INTEGER    ,
  cart_term_description  TEXT       ,
  cart_term_name         TEXT       ,
  cart_term_duration     INTEGER    ,
  shopping_cart_id       INTEGER     NOT NULL,
  product_offering_id    INTEGER     NOT NULL,
  item_price             FLOAT,
  item_totalprice        FLOAT,
  item_price_id          INTEGER     NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE Product
(
  id                     INTEGER     NOT NULL,
  href                   TEXT       ,
  description            TEXT       ,
  status                 TEXT       ,
  isBundle               BOOLEAN    ,
  isCustomerVisible      BOOLEAN    ,
  lastUpdate             TIMESTAMP  ,
  name                   TEXT       ,
  product_serial_number  TEXT       ,
  product_term_description  TEXT    ,
  product_term_name      TEXT       ,
  product_term_duration  INTEGER    ,
  orderDateTime          DATE       ,
  startDateTime          DATE       ,
  terminationDateTime    DATE       ,
  product_price          FLOAT,
  product_specification_id       INTEGER     NOT NULL,
  product_price_id               INTEGER     NOT NULL,
  product_relatedParty_id        INTEGER     NOT NULL,
  agreemnet_item_id              INTEGER     NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE ProductOffering
(
  id                     INTEGER     NOT NULL,
  href                   TEXT       ,
  description            TEXT       ,
  lifecycleStatus        TEXT       ,
  isBundle               BOOLEAN    ,
  lastUpdate             TIMESTAMP  ,
  name                   TEXT       ,
  statusReason           TEXT       ,
  version                VARCHAR(20),
  isSellable             BOOLEAN    ,
  validFor_startDateTime DATE       ,
  validFor_endDateTime   DATE       ,
  product_id             INTEGER     NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE ContactMedium
(
  id                     INTEGER     NOT NULL,
  contact_type           TEXT       ,
  phone_number           TEXT       ,
  email_address          TEXT       ,
  preffred               BOOLEAN    ,
  validFor_startDateTime DATE       ,
  validFor_endDateTime   DATE       ,
  PRIMARY KEY (id)
);

CREATE TABLE AgreemnetItem
(
  id                     INTEGER     NOT NULL,
  href                   TEXT       ,
  name                   TEXT       ,
  PRIMARY KEY (id)
);

CREATE TABLE CartPrice
(
  id                   INTEGER     NOT NULL,
  description          TEXT     ,
  name                 TEXT     ,
  pricetype            TEXT     ,
  unitofmeasure        TEXT     ,
  recuringchargeperiod text     ,
  product_price_id     INTEGER     NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE ProductPrice
(
  id                   INTEGER     NOT NULL,
  percentage           NUMERIC    ,
  taxRate              NUMERIC    ,
  dutyFreeAmount_id    INTEGER     NOT NULL,
  taxIncludedAmount_id INTEGER     NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE Money
(
  id    INTEGER     NOT NULL,
  unit  VARCHAR(10),
  value NUMERIC    ,
  PRIMARY KEY (id)
);

CREATE TABLE ProductSpecification
(
  id              INTEGER     NOT NULL,
  href            TEXT       ,
  name            TEXT       ,
  description     TEXT       ,
  brand           TEXT       ,
  isBundle        BOOLEAN    ,
  lastUpdate      TIMESTAMP  ,
  lifecycleStatus VARCHAR(20),
  productNumber   VARCHAR(50),
  version         VARCHAR(20),
  PRIMARY KEY (id)
);

CREATE TABLE RelatedParty
(
  id   INTEGER     NOT NULL,
  href TEXT       ,
  name TEXT       ,
  role VARCHAR(50),
  PRIMARY KEY (id)
);


ALTER TABLE ShoppingCart
  ADD CONSTRAINT FK_ShoppingCart_TO_CartPrice
    FOREIGN KEY (cart_price_id)
    REFERENCES CartPrice (id);

ALTER TABLE ShoppingCart
  ADD CONSTRAINT FK_ShoppingCart_TO_ContactMedium
    FOREIGN KEY (contact_medium_id)
    REFERENCES ContactMedium (id);

ALTER TABLE ShoppingCart
  ADD CONSTRAINT FK_ShoppingCart_TO_RelatedParty
    FOREIGN KEY (product_relatedParty_id)
    REFERENCES RelatedParty (id);


ALTER TABLE CartItem
  ADD CONSTRAINT FK_CartItem_TO_ShoppingCart
    FOREIGN KEY (shopping_cart_id)
    REFERENCES ShoppingCart (id);

ALTER TABLE CartItem
  ADD CONSTRAINT FK_CartItem_TO_CartPrice
    FOREIGN KEY (item_price_id)
    REFERENCES  CartPrice (id);

ALTER TABLE CartItem
  ADD CONSTRAINT FK_CartItem_TO_ProductOffering
    FOREIGN KEY (product_offering_id)
    REFERENCES  ProductOffering (id);

ALTER TABLE ProductOffering
  ADD CONSTRAINT FK_ProductOffering_TO_Product
    FOREIGN KEY (product_id)
    REFERENCES  Product (id);

ALTER TABLE Product
  ADD CONSTRAINT FK_Product_TO_ProductSpecification
    FOREIGN KEY (product_specification_id)
    REFERENCES  ProductSpecification (id);

ALTER TABLE Product
  ADD CONSTRAINT FK_Product_TO_RelatedParty
    FOREIGN KEY (product_relatedParty_id)
    REFERENCES  RelatedParty (id);

ALTER TABLE Product
  ADD CONSTRAINT FK_Product_TO_AgreemnetItem
    FOREIGN KEY (agreemnet_item_id)
    REFERENCES  AgreemnetItem (id);

ALTER TABLE Product
  ADD CONSTRAINT FK_Product_TO_ProductPrice
    FOREIGN KEY (product_price_id)
    REFERENCES  ProductPrice (id);
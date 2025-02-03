-- +migrate Up
-- +migrate StatementBegin

-- Create ENUM type for payment_type
CREATE TYPE payment_type_enum AS ENUM ('CASH', 'CREDIT');

-- Create table invoices
CREATE TABLE invoices (
    id SERIAL PRIMARY KEY,                                     -- Auto-incremented unique identifier
    invoice_no TEXT NOT NULL UNIQUE,                            -- Invoice number (required: true, type: text, minLength: 1)
    date DATE NOT NULL,                                         -- Date (required: true, type: date)
    customer_name TEXT NOT NULL CHECK (LENGTH(customer_name) >= 2), -- Customer name (required: true, type: text, minLength: 2)
    salesperson_name TEXT NOT NULL CHECK (LENGTH(salesperson_name) >= 2), -- Salesperson name (required: true, type: text, minLength: 2)
    payment_type payment_type_enum NOT NULL,                   -- Payment type (required: true, type: ENUM, values: "CASH" | "CREDIT")
    notes TEXT,                                                -- Notes for additional information (optional, type: text)
    CONSTRAINT chk_notes_length CHECK (LENGTH(notes) >= 5)     -- Enforces that notes, if provided, are at least 5 characters long
);

-- Create table products
CREATE TABLE products (
    id SERIAL PRIMARY KEY,                                      -- Product ID (auto-incremented)
    invoice_no TEXT NOT NULL,                                    -- Invoice number (required: true, type: text)
    item_name TEXT NOT NULL CHECK (LENGTH(item_name) >= 5),       -- Item name (required: true, type: text, minLength: 5)
    quantity INT NOT NULL CHECK (quantity >= 1),                  -- Quantity (required: true, type: number, minValue: 1)
    total_cost NUMERIC(10, 2) NOT NULL CHECK (total_cost >= 0),   -- Total cost of goods sold (required: true, type: number, minValue: 0)
    total_price NUMERIC(10, 2) NOT NULL CHECK (total_price >= 0), -- Total price sold (required: true, type: number, minValue: 0)
    FOREIGN KEY (invoice_no) REFERENCES invoices(invoice_no) ON DELETE CASCADE  -- Foreign key reference to invoices (deletes products when invoice is deleted)
);

-- +migrate StatementEnd

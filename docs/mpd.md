# Modèle Physique des Données (MPD)

The MPD represents the technical translation of the MLD into the SQL language
## 🗄️ Tables SQL

```sql
CREATE TABLE user (
    id_user SERIAL PRIMARY KEY,
    name_user VARCHAR(50) NOT NULL,
    email_user VARCHAR(100) UNIQUE NOT NULL,
    password_user VARCHAR(255) NOT NULL,
    created_at_user TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE category (
    id_category SERIAL PRIMARY KEY,
    nom_category VARCHAR(100) NOT NULL,
    description_category TEXT
);

CREATE TABLE expense (
    id_expense SERIAL PRIMARY KEY,
    description_expense TEXT,
    amount DECIMAL(10, 2) NOT NULL,
    date_expense DATE NOT NULL,
    created_at_expense TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at_expense TIMESTAMP, 
    id_user INT NOT NULL,
    id_category INT,
    FOREIGN KEY (id_user) REFERENCES user(id_user),
    FOREIGN KEY (id_category) REFERENCES category(id_category)
);
```
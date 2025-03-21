# Modèle Logique des Données (MLD)

MLD describes business entities and their relationships in an abstract manner, without regard for technical constraints.

User (
    id_user (PK),
    email_user (UNIQUE),
    password_user,
    created_at_user,
    updated_at_user
)

Expense (
    id_expense (PK),
    description,
    amount_expense,
    created_at_expense,
    updated_at_expense,
    date_expense,
    id_user (FK -> User),
    id_category (FK -> Category)
)

Category (
    id_category (PK),
    name_category (UNIQUE)
)

## 📊 Entités principales et attributs
- **Utilisateur**
  - ID utilisateur (clé primaire)
  - Nom d'utilisateur
  - Email
  - Mot de passe (haché)
  - Date de création
- **Catégorie de dépenses**
  - ID catégorie (clé primaire)
  - Nom de la catégorie
  - Description
- **Dépense**
  - ID dépense (clé primaire)
  - Montant
  - Date de la dépense
  - Description
  - ID utilisateur (clé étrangère)
  - ID catégorie (clé étrangère)

## 🔗 Relations entre les entités
- Un utilisateur peut créer plusieurs dépenses.
- Une dépense appartient à une catégorie.
- Une catégorie peut être partagée par plusieurs dépenses.

## 💡 Hypothèses et choix de modélisation
- Les utilisateurs ne peuvent voir que leurs propres dépenses.
- Les catégories sont personnalisables par utilisateur.
- Les montants sont stockés en nombre décimal(deux chiffres après la virgule) pour la précision.

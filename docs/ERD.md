
# Gym Management System — Entity Relationship Diagram

## Database Overview

Database ini menggunakan PostgreSQL dan terdiri dari 8 tabel utama:

1. users
2. trainers
3. members
4. membership_plans
5. memberships
6. payments
7. training_schedules
8. attendances

## Entity Relationship Diagram

```mermaid
erDiagram
    USERS ||--o| TRAINERS : "has profile"
    MEMBERS ||--o{ MEMBERSHIPS : "owns"
    MEMBERSHIP_PLANS ||--o{ MEMBERSHIPS : "used by"
    MEMBERSHIPS ||--o| PAYMENTS : "has payment"
    MEMBERS ||--o{ TRAINING_SCHEDULES : "books"
    TRAINERS ||--o{ TRAINING_SCHEDULES : "handles"
    MEMBERS ||--o{ ATTENDANCES : "records"

    USERS {
        uuid id PK
        varchar username UK
        varchar full_name
        text password_hash
        varchar role
        timestamptz created_at
        timestamptz updated_at
    }

    TRAINERS {
        uuid id PK
        uuid user_id FK,UK
        varchar specialization
        varchar phone
        timestamptz created_at
        timestamptz updated_at
    }

    MEMBERS {
        uuid id PK
        varchar name
        varchar email UK
        varchar phone
        text address
        timestamptz created_at
        timestamptz updated_at
    }

    MEMBERSHIP_PLANS {
        uuid id PK
        varchar name
        bigint price
        int duration_days
        text description
        boolean is_active
        timestamptz created_at
        timestamptz updated_at
    }

    MEMBERSHIPS {
        uuid id PK
        uuid member_id FK
        uuid plan_id FK
        date start_date
        date end_date
        varchar status
        timestamptz created_at
        timestamptz updated_at
    }

    PAYMENTS {
        uuid id PK
        uuid membership_id FK,UK
        bigint amount
        varchar payment_method
        varchar status
        timestamptz paid_at
        timestamptz created_at
    }

    TRAINING_SCHEDULES {
        uuid id PK
        uuid member_id FK
        uuid trainer_id FK
        timestamptz start_at
        timestamptz end_at
        varchar status
        text notes
        timestamptz created_at
        timestamptz updated_at
    }

    ATTENDANCES {
        uuid id PK
        uuid member_id FK
        timestamptz check_in
        timestamptz check_out
        timestamptz created_at
    }
```

## Relationship Description

| Parent Table | Child Table | Relationship |
|---|---|---|
| users | trainers | One-to-zero-or-one |
| members | memberships | One-to-many |
| membership_plans | memberships | One-to-many |
| memberships | payments | One-to-zero-or-one |
| members | training_schedules | One-to-many |
| trainers | training_schedules | One-to-many |
| members | attendances | One-to-many |

## Important Database Rules

- A trainer account is stored in `users` and has a corresponding record in `trainers`.
- A membership belongs to one member and uses one membership plan.
- Each membership can have at most one payment record.
- A training schedule belongs to one member and one trainer.
- An attendance record belongs to one member.
- Passwords are stored as bcrypt hashes, never as plain text.
- Membership activation after successful payment is handled by the service layer.
- Trainer schedule conflict validation is handled by the service layer.
- Only active members should be allowed to check in, enforced by business logic.

## Notes

This document describes the initial database schema implemented in the `feature/project-setup` branch.

Any schema changes must be coordinated with the project team and reflected in the migration files.

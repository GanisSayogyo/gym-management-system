-- ==========================================
-- INITIAL SEED: USERS
-- ==========================================

INSERT INTO users (id, username, full_name, password_hash, role)
VALUES
(
    '11111111-1111-4111-8111-111111111111',
    'admin',
    'Administrator Gym',
    '$2a$10$VbDxiyaYfDlGSAEwbyQBXeaC0M18Jc2XJWcGqTdJwq223VmN8bBw2',
    'admin'
),
(
    '22222222-2222-4222-8222-222222222222',
    'trainer.andri',
    'Andri Pratama',
    '$2a$10$Zg2/NRE3mH7W96lD.tJMHOBOm0rPFZUf.Qz/AfKSwLwnNR0S9lWy2',
    'trainer'
),
(
    '33333333-3333-4333-8333-333333333333',
    'trainer.sinta',
    'Sinta Maharani',
    '$2a$10$Zg2/NRE3mH7W96lD.tJMHOBOm0rPFZUf.Qz/AfKSwLwnNR0S9lWy2',
    'trainer'
)
ON CONFLICT (id) DO NOTHING;

-- ==========================================
-- INITIAL SEED: TRAINERS
-- ==========================================

INSERT INTO trainers (id, user_id, specialization, phone)
VALUES
(
    '44444444-4444-4444-8444-444444444444',
    '22222222-2222-4222-8222-222222222222',
    'Strength Training',
    '081234567801'
),
(
    '55555555-5555-4555-8555-555555555555',
    '33333333-3333-4333-8333-333333333333',
    'Fitness and Cardio',
    '081234567802'
)
ON CONFLICT (id) DO NOTHING;

-- ==========================================
-- INITIAL SEED: MEMBERS
-- ==========================================

INSERT INTO members (id, name, email, phone, address)
VALUES
(
    '66666666-6666-4666-8666-666666666666',
    'Budi Santoso',
    'budi@example.com',
    '081234567803',
    'Surabaya'
),
(
    '77777777-7777-4777-8777-777777777777',
    'Citra Permata',
    'citra@example.com',
    '081234567804',
    'Sidoarjo'
),
(
    '88888888-8888-4888-8888-888888888888',
    'Dimas Saputra',
    'dimas@example.com',
    '081234567805',
    'Gresik'
)
ON CONFLICT (id) DO NOTHING;

-- ==========================================
-- INITIAL SEED: MEMBERSHIP PLANS
-- ==========================================

INSERT INTO membership_plans (id, name, price, duration_days, description, is_active)
VALUES
(
    '99999999-9999-4999-8999-999999999991',
    'Basic Monthly',
    250000,
    30,
    'Gym access for 30 days',
    TRUE
),
(
    '99999999-9999-4999-8999-999999999992',
    'Standard Quarterly',
    650000,
    90,
    'Gym access for 90 days',
    TRUE
),
(
    '99999999-9999-4999-8999-999999999993',
    'Premium Annual',
    2000000,
    365,
    'Gym access for 365 days',
    TRUE
)
ON CONFLICT (id) DO NOTHING;

-- ==========================================
-- INITIAL SEED: MEMBERSHIPS
-- ==========================================

INSERT INTO memberships (
    id, member_id, plan_id, start_date, end_date, status
)
VALUES
(
    'aaaaaaaa-aaaa-4aaa-8aaa-aaaaaaaaaaa1',
    '66666666-6666-4666-8666-666666666666',
    '99999999-9999-4999-8999-999999999991',
    CURRENT_DATE - 10,
    CURRENT_DATE + 20,
    'active'
),
(
    'aaaaaaaa-aaaa-4aaa-8aaa-aaaaaaaaaaa2',
    '77777777-7777-4777-8777-777777777777',
    '99999999-9999-4999-8999-999999999992',
    CURRENT_DATE - 20,
    CURRENT_DATE + 70,
    'active'
),
(
    'aaaaaaaa-aaaa-4aaa-8aaa-aaaaaaaaaaa3',
    '88888888-8888-4888-8888-888888888888',
    '99999999-9999-4999-8999-999999999993',
    CURRENT_DATE,
    CURRENT_DATE + 365,
    'pending'
)
ON CONFLICT (id) DO NOTHING;

-- ==========================================
-- INITIAL SEED: PAYMENTS
-- ==========================================

INSERT INTO payments (
    id, membership_id, amount, payment_method, status, paid_at
)
VALUES
(
    'bbbbbbbb-bbbb-4bbb-8bbb-bbbbbbbbbbb1',
    'aaaaaaaa-aaaa-4aaa-8aaa-aaaaaaaaaaa1',
    250000,
    'cash',
    'paid',
    NOW()
),
(
    'bbbbbbbb-bbbb-4bbb-8bbb-bbbbbbbbbbb2',
    'aaaaaaaa-aaaa-4aaa-8aaa-aaaaaaaaaaa2',
    650000,
    'transfer',
    'paid',
    NOW()
),
(
    'bbbbbbbb-bbbb-4bbb-8bbb-bbbbbbbbbbb3',
    'aaaaaaaa-aaaa-4aaa-8aaa-aaaaaaaaaaa3',
    2000000,
    'cash',
    'pending',
    NULL
)
ON CONFLICT (id) DO NOTHING;

-- ==========================================
-- INITIAL SEED: TRAINING SCHEDULES
-- ==========================================

INSERT INTO training_schedules (
    id, member_id, trainer_id, start_at, end_at, status, notes
)
VALUES
(
    'cccccccc-cccc-4ccc-8ccc-ccccccccccc1',
    '66666666-6666-4666-8666-666666666666',
    '44444444-4444-4444-8444-444444444444',
    NOW() - INTERVAL '2 days',
    NOW() - INTERVAL '2 days' + INTERVAL '1 hour',
    'completed',
    'Strength training session'
),
(
    'cccccccc-cccc-4ccc-8ccc-ccccccccccc2',
    '77777777-7777-4777-8777-777777777777',
    '55555555-5555-4555-8555-555555555555',
    NOW() + INTERVAL '1 day',
    NOW() + INTERVAL '1 day' + INTERVAL '1 hour',
    'scheduled',
    'Cardio training session'
)
ON CONFLICT (id) DO NOTHING;

-- ==========================================
-- INITIAL SEED: ATTENDANCES
-- ==========================================

INSERT INTO attendances (
    id, member_id, check_in, check_out
)
VALUES
(
    'dddddddd-dddd-4ddd-8ddd-ddddddddddd1',
    '66666666-6666-4666-8666-666666666666',
    NOW() - INTERVAL '1 day',
    NOW() - INTERVAL '1 day' + INTERVAL '2 hours'
),
(
    'dddddddd-dddd-4ddd-8ddd-ddddddddddd2',
    '77777777-7777-4777-8777-777777777777',
    NOW() - INTERVAL '1 hour',
    NULL
)
ON CONFLICT (id) DO NOTHING;


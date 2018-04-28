INSERT INTO workouts (workout_id, user_id, num_ex, workout_name, date_complete) 
    VALUES 
    ('we32ad6e2-5957-4e29-a14f-6118bc686f26', 'u00d9e2e9-40b3-4395-9b79-23be002719e0', 2, 'New Workout', '2018-04-27'),
    ('w2123155c-f442-451a-9916-6cb6b98894a2', 'u00d9e2e9-40b3-4395-9b79-23be002719e0', 1, 'Old Workout', '2016-05-21'),
    ('wf69e747a-df96-4f4a-8426-5bc2bf1be6b0', 'u00d9e2e9-40b3-4395-9b79-23be002719e0', 0, 'Middle Workout', '2017-08-02');

INSERT INTO exercises (ex_id, workout_id, ex_order, num_sets, ex_name)
    VALUES
    ('eb50c1798-75b3-489b-a352-ff151002868f', 'we32ad6e2-5957-4e29-a14f-6118bc686f26', 0, 2, 'Bench Press'),
    ('e43d6ecc0-2036-4fd1-a483-2bf09fae175c', 'we32ad6e2-5957-4e29-a14f-6118bc686f26', 1, 0, 'Shoulder Press'),
    ('ebe3be9be-6596-4f0d-8645-b499d3c6c7c1', 'w2123155c-f442-451a-9916-6cb6b98894a2', 0, 1, 'Weighted Dips');

INSERT INTO exercise_sets (set_id, ex_id, set_order, reps, set_weight, note)
    VALUES
    ('s12cad48b-5e5c-458e-8f2f-a95d6be8ec2a', 'eb50c1798-75b3-489b-a352-ff151002868f', 0, 8, 170.00, 'Awesome'),
    ('sca1bbcde-1223-4ee2-a337-c9e2fa87fdc7', 'eb50c1798-75b3-489b-a352-ff151002868f', 1, 12, 130, 'Dropped Weight'),
    ('s5fd135a3-e978-47e4-aa92-a8e2b7f1c0c8', 'ebe3be9be-6596-4f0d-8645-b499d3c6c7c1', 0, 8, 45, 'Last Workout Dippin');
-- Inserting data into 'credential' table
INSERT INTO credential (email, password) VALUES 
('john.doe@example.com', 'password123'),
('jane.smith@example.com', 'securepass'),
('bob.jones@example.com', 'password456');

-- Inserting data into 'user' table
INSERT INTO users (credential_id, name, age, gender) VALUES
(1, 'John Doe', 28, 'Male'),
(2, 'Jane Smith', 32, 'Female'),
(3, 'Bob Jones', 24, 'Male');

-- Inserting data into 'todo' table
INSERT INTO todo (user_id, title, description, done) VALUES
(1, 'Buy groceries', 'Buy milk, bread, and eggs', 'false'),
(1, 'Workout', 'Go to the gym for an hour', 'true'),
(2, 'Read book', 'Finish reading the novel', 'false'),
(2, 'Pay bills', 'Pay electricity and water bills', 'true'),
(3, 'Clean house', 'Vacuum and dust the living room', 'false');

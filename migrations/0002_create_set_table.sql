CREATE TABLE sets (
    id INT AUTO_INCREMENT PRIMARY KEY,
    reps INT,
    weight DECIMAL(10,2),
    exercise INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (exercise) REFERENCES exercises(id)
);
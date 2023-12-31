CREATE TABLE IF NOT EXISTS Users(
    user_id INT PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS Segments(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS UserSegment(
    id BIGSERIAL PRIMARY KEY,
    user_id INT REFERENCES Users(user_id) ON DELETE CASCADE ,
    name VARCHAR REFERENCES Segments(name) ON DELETE CASCADE
);
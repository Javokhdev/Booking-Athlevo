
CREATE TABLE facilities (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(100),
    image VARCHAR(255),
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT
);

CREATE TABLE sport_halls (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    location VARCHAR(255) NOT NULL,
    longitude VARCHAR(50),
    latitude VARCHAR(50),
    type_sport VARCHAR(100),
    type_gender VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT
);

CREATE TABLE subscription_personal (
    id UUID PRIMARY KEY,
    gym_id UUID REFERENCES sport_halls(id),
    type VARCHAR(100),
    description TEXT,
    price INT,
    duration INT,
    count INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT
);

CREATE TABLE subscription_group (
    id UUID PRIMARY KEY,
    gym_id UUID REFERENCES sport_halls(id),
    coach_id UUID, -- REFERENCES users(id),
    type VARCHAR(100),
    description TEXT,
    price INT,
    capacity INT,
    time TIMESTAMP,
    duration INT,
    count INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT
);

CREATE TABLE subscription_coach (
    id UUID PRIMARY KEY,
    gym_id UUID REFERENCES sport_halls(id),
    coach_id UUID, -- REFERENCES users(id),
    type VARCHAR(100),
    description TEXT,
    price INT,
    duration INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT
);

CREATE TABLE booking_personal (
    id UUID PRIMARY KEY,
    user_id UUID, -- REFERENCES users(id),
    subscription_id UUID REFERENCES subscription_personal(id),
    payment INT,
    access_status VARCHAR(100),
    start_date TIMESTAMP,
    count INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT
);


CREATE TABLE access_personal (
    booking_id UUID REFERENCES booking_personal(id),
    date TIMESTAMP
);

CREATE TABLE booking_group (
    id UUID PRIMARY KEY,
    user_id UUID ,-- REFERENCES users(id),
    subscription_id UUID REFERENCES subscription_group(id),
    payment INT,
    access_status VARCHAR(100),
    start_date TIMESTAMP,
    count INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT
);

CREATE TABLE access_group (
    booking_id UUID REFERENCES booking_group(id),
    date TIMESTAMP
);

CREATE TABLE booking_coach (
    id UUID PRIMARY KEY,
    user_id UUID ,--REFERENCES users(id),
    subscription_id UUID  REFERENCES subscription_coach(id),
    payment INT,
    access_status VARCHAR(100),
    start_date TIMESTAMP,
    count INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT
);

CREATE TABLE access_coach (
    booking_id UUID REFERENCES booking_coach(id),
    date TIMESTAMP
);

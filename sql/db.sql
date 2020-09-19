
CREATE TABLE institute
(
  ID serial NOT NULL,
  name VARCHAR(128) NOT NULL,
  PRIMARY KEY (ID)
);

CREATE TABLE patient
(
  ID serial NOT NULL,
  name VARCHAR(16) NOT NULL,
  age INT NOT NULL,
  gender INT NOT NULL,
  PRIMARY KEY (ID)
);

CREATE TABLE doctor
(
  ID VARCHAR(16) NOT NULL,
  password VARCHAR(255) NOT NULL,
  master bool NOT NULL,
  instituteID serial NOT NULL,
  PRIMARY KEY (ID),
  FOREIGN KEY (instituteID) REFERENCES institute(ID) on delete cascade on update cascade
);

CREATE TABLE log
(
  ID serial NOT NULL,
  exp json NOT NULL,
  work json NOT NULL,
  chk json NOT NULL,
  date VARCHAR(32) NOT NULL,
  doctorID VARCHAR(16) NOT NULL,
  patientID serial NOT NULL,
  PRIMARY KEY (ID),
  FOREIGN KEY (doctorID) REFERENCES doctor(ID) on delete cascade on update cascade,
  FOREIGN KEY (patientID) REFERENCES patient(ID) on delete cascade on update cascade
);
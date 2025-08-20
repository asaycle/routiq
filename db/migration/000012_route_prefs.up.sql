CREATE TABLE route_prefs (
  route_id CHAR(20) NOT NULL,
  pref_id  CHAR(6)  NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  PRIMARY KEY (route_id, pref_id),
  CONSTRAINT fk_route_prefs_route
    FOREIGN KEY (route_id) REFERENCES routes(id) ON DELETE CASCADE,
  CONSTRAINT fk_route_prefs_pref
    FOREIGN KEY (pref_id)  REFERENCES prefs(id)  ON DELETE RESTRICT
);

CREATE INDEX idx_route_prefs_route ON route_prefs(route_id);
CREATE INDEX idx_route_prefs_pref  ON route_prefs(pref_id);
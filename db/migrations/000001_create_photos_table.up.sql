CREATE OR REPLACE FUNCTION update_updated_at()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

create table if not exists photos(
  id uuid primary key default gen_random_uuid(),
  location text not null,
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now()
);

CREATE TRIGGER photos_update_updated_at
BEFORE UPDATE ON photos
FOR EACH ROW
EXECUTE FUNCTION update_updated_at();

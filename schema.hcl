schema "public" {
  comment = "Default EVEBot Tools schema"
}

table "server_status" {
  schema = schema.public
  column "id" {
    type = bigint
  }
  column "players" {
    type = int
  }
  column "server_version" {
    type = text
  }
  column "start_time" {
    type = timestamptz
  }
  column "vip" {
    type = bool
  }
  primary_key {
    columns = [column.id]
  }
}

table "killmail" {
  schema = schema.public
  column "id" {
    type = bigint
  }
  column "killmail_time" {
    type = timestamptz
  }
  column "solar_system_id" {
    type = int
  }
  column "war_id" {
    type = bigint
    null = true
  }
  column "damage_taken" {
    type = int
  }
  column "total_value" {
    type = float
  }

  primary_key {
    columns = [column.id]
  }
}

table "killmail_victim" {
  schema = schema.public
  column "id" {
    type = bigint
  }
  column "character_id" {
    type = int
  }
  column "corporation_id" {
    type = int
  }
  column "alliance_id" {
    type = int
    null = true
  }
  column "faction_id" {
    type = int
    null = true
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "killmail_id" {
    columns     = [column.id]
    ref_columns = [table.killmail.column.id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
}

table "killmail_attacker" {
  schema = schema.public
  column "id" {
    type = bigint
  }
  column "character_id" {
    type = int
    null = true
  }
  column "corporation_id" {
    type = int
    null = true
  }
  column "alliance_id" {
    type = int
    null = true
  }
  column "faction_id" {
    type = int
    null = true
  }
  column "damage_done" {
    type = int
  }
  column "final_blow" {
    type = bool
  }
  column "security_status" {
    type = float
  }
  column "weapon_type_id" {
    type = int
  }
  column "ship_type_id" {
    type = int
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "killmail_id" {
    columns     = [column.id]
    ref_columns = [table.killmail.column.id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
}

table "killmail_item" {
  schema = schema.public
  column "id" {
    type = bigint
  }
  column "type_id" {
    type = int
  }
  column "flag" {
    type = int
  }
  column "quantity_destroyed" {
    type = int
    null = true
  }
  column "quantity_dropped" {
    type = int
    null = true
  }
  column "singleton" {
    type = bool
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "killmail_id" {
    columns     = [column.id]
    ref_columns = [table.killmail.column.id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
}

table "killmail_stats" {
  schema = schema.public
  column "id" {
    type = bigint
  }
  column "location_id" {
    type = int
  }
  column "hash" {
    type = text
  }
  column "fitted_value" {
    type = float
  }
  column "destroyed_value" {
    type = float
  }
  column "dropped_value" {
    type = float
  }
  column "href" {
    type = text
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "killmail_id" {
    columns     = [column.id]
    ref_columns = [table.killmail.column.id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
}

table "universe_type" {
  schema = schema.public
  column "id" {
    type = int
  }
  column "name" {
    type = text
  }
  column "description" {
    type = text
  }
  column "capacity" {
    type = float
    null = true
  }
  column "graphic_id" {
    type = int
    null = true
  }
  column "group_id" {
    type = int
  }
  column "icon_id" {
    type = int
    null = true
  }
  column "market_group_id" {
    type = int
    null = true
  }
  column "mass" {
    type = float
    null = true
  }
  column "packaged_volume" {
    type = float
    null = true
  }
  column "portion_size" {
    type = int
    null = true
  }
  column "published" {
    type = bool
  }
  column "radius" {
    type = float
    null = true
  }
  column "volume" {
    type = float
    null = true
  }

  primary_key {
    columns = [column.id]
  }
}

table "universe_type_dogma_attribute" {
  schema = schema.public
  column "type_id" {
    type = int
  }
  column "attribute_id" {
    type = int
  }
  column "value" {
    type = float
  }
  primary_key {
    columns = [column.type_id, column.attribute_id]
  }
  foreign_key "type_id" {
    columns     = [column.type_id]
    ref_columns = [table.universe_type.column.id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
}

table "universe_type_dogma_effect" {
  schema = schema.public
  column "type_id" {
    type = int
  }
  column "effect_id" {
    type = int
  }
  column "is_default" {
    type = bool
  }
  primary_key {
    columns = [column.type_id, column.effect_id]
  }
  foreign_key "type_id" {
    columns     = [column.type_id]
    ref_columns = [table.universe_type.column.id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
}

table "universe_system" {
  schema = schema.public
  column "id" {
    type = int
  }
  column "name" {
    type = text
  }
  column "security_class" {
    type = text

  }
  column "security_status" {
    type = float
    null = true
  }
  column "constellation_id" {
    type = int
    null = true
  }
  column "star_id" {
    type = int
    null = true
  }
  primary_key {
    columns = [column.id]
  }

}

table "universe_system_position" {
  schema = schema.public
  column "id" {
    type = int
  }
  column "position_x" {
    type = float
    null = true
  }
  column "position_y" {
    type = float
    null = true
  }
  column "position_z" {
    type = float
    null = true
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "system_id" {
    columns     = [column.id]
    ref_columns = [table.universe_system.column.id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
}

table "universe_planet" {
  schema = schema.public
  column "id" {
    type = int
  }
  column "system_id" {
    type = int
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "system_id" {
    columns     = [column.system_id]
    ref_columns = [table.universe_system.column.id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
}

table "universe_moon" {
  schema = schema.public
  column "id" {
    type = int
  }
  column "system_id" {
    type = int
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "system" {
    columns     = [column.system_id]
    ref_columns = [table.universe_system.column.id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
}

table "universe_astroid_belt" {
  schema = schema.public
  column "id" {
    type = int
  }
  column "system_id" {
    type = int
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "system_id" {
    columns     = [column.system_id]
    ref_columns = [table.universe_system.column.id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
}
schema "public" {
  comment = "Default EVEBot Tools schema"
}

table "server_status" {
  schema = schema.public
  column "id" {
    type = bigserial
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
    type = bigserial
  }
  column "killmail_id" {
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
    null = true
  }
  column "ship_type_id" {
    type = int
    null = true
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "killmail_id" {
    columns     = [column.killmail_id]
    ref_columns = [table.killmail.column.id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
}

table "killmail_item" {
  schema = schema.public
  column "id" {
    type = bigserial
  }
  column "killmail_id" {
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
    columns     = [column.killmail_id]
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
  }
  column "packaged_volume" {
    type = float
  }
  column "portion_size" {
    type = int
  }
  column "published" {
    type = bool
  }
  column "radius" {
    type = float
  }
  column "volume" {
    type = float
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
  }
  column "constellation_id" {
    type = int
  }
  column "star_id" {
    type = int
  }
  column "position_x" {
    type = float
  }
  column "position_y" {
    type = float
  }
  column "position_z" {
    type = float
  }
  primary_key {
    columns = [column.id]
  }
#  foreign_key "constellation_id" {
#    columns     = [column.constellation_id]
#    ref_columns = [table.universe_constellation.column.id]
#    on_update   = CASCADE
#    on_delete   = CASCADE
#  }

}

table "universe_planet" {
  schema = schema.public
  column "id" {
    type = int
  }
  column "system_id" {
    type = int
  }
  column "name" {
    type = text
  }
  column "position_x" {
    type = float
  }
  column "position_y" {
    type = float
  }
  column "position_z" {
    type = float
  }
  column "type_id" {
    type = int
  }
  primary_key {
    columns = [column.id]
  }
#  foreign_key "system_id" {
#    columns     = [column.system_id]
#    ref_columns = [table.universe_system.column.id]
#    on_update   = CASCADE
#    on_delete   = CASCADE
#  }
#  foreign_key "type_id" {
#    columns     = [column.type_id]
#    ref_columns = [table.universe_type.column.id]
#    on_update   = CASCADE
#    on_delete   = CASCADE
#  }
}

table "universe_moon" {
  schema = schema.public
  column "id" {
    type = int
  }
  column "system_id" {
    type = int
  }
  column "name" {
    type = text
  }
  column "position_x" {
    type = float
  }
  column "position_y" {
    type = float
  }
  column "position_z" {
    type = float
  }
  primary_key {
    columns = [column.id]
  }
#  foreign_key "system" {
#    columns     = [column.system_id]
#    ref_columns = [table.universe_system.column.id]
#    on_update   = CASCADE
#    on_delete   = CASCADE
#  }
}

table "universe_astroid_belt" {
  schema = schema.public
  column "id" {
    type = int
  }
  column "name" {
    type = text
  }
  column "system_id" {
    type = int
  }
  column "position_x" {
    type = float
  }
  column "position_y" {
    type = float
  }
  column "position_z" {
    type = float
  }
  primary_key {
    columns = [column.id]
  }
#  foreign_key "system_id" {
#    columns     = [column.system_id]
#    ref_columns = [table.universe_system.column.id]
#    on_update   = CASCADE
#    on_delete   = CASCADE
#  }
}

table "universe_ancestry" {
  schema = schema.public
  column "id" {
    type = int
  }
  column "name" {
    type = text
  }
  column "bloodline_id" {
    type = int
  }
  column "description" {
    type = text
  }
  column "icon_id" {
    type = int
    null = true
  }
  column "short_description" {
    type = text
    null = true
  }
  primary_key {
    columns = [column.id]
  }
}

table "universe_bloodlines" {
  schema = schema.public
  column "id" {
    type = int
  }
  column "corporation_id" {
    type = int
  }
  column "description" {
    type = text
  }
  column "race_id" {
    type = int
  }
  column "ship_type_id" {
    type = int
  }
  column "charisma" {
    type = int
  }
  column "intelligence" {
    type = int
  }
  column "memory" {
    type = int
  }
  column "perception" {
    type = int
  }
  column "willpower" {
    type = int
  }
  primary_key {
    columns = [column.id]
  }
}

table "universe_category" {
  schema = schema.public
  column "id" {
    type = int
  }
  column "name" {
    type = text
  }
  column "published" {
    type = bool
  }
  primary_key {
    columns = [column.id]
  }
}

table "universe_constellation" {
  schema = schema.public
  column "id" {
    type = int
  }
  column "name" {
    type = text
  }
  column "region_id" {
    type = int
  }
  column "position_x" {
    type = float
  }
  column "position_y" {
    type = float
  }
  column "position_z" {
    type = float
  }
  primary_key {
    columns = [column.id]
  }
#  foreign_key "region_id" {
#    columns     = [column.region_id]
#    ref_columns = [table.universe_region.column.id]
#    on_update   = CASCADE
#    on_delete   = CASCADE
#  }
}

table "universe_faction" {
  schema = schema.public
  column "id" {
    type = int
  }
  column "corporation_id" {
    type = int
    null = true
  }
  column "description" {
    type = text
  }
  column "is_unique" {
    type = bool
  }
  column "militia_corporation_id" {
    type = int
    null = true
  }
  column "name" {
    type = text
  }
  column "size_factor" {
    type = float
  }
  column "solar_system_id" {
    type = int
    null = true
  }
  column "station_count" {
    type = int
  }
  column "station_system_count" {
    type = int
  }
  primary_key {
    columns = [column.id]
  }
  #  foreign_key "corporation_id" {
  #    columns     = [column.corporation_id]
  #    ref_columns = [table.universe_corporation.column.id]
  #    on_update   = CASCADE
  #    on_delete   = CASCADE
  #  }
  #  foreign_key "militia_corporation_id" {
  #    columns     = [column.militia_corporation_id]
  #    ref_columns = [table.universe_corporation.column.id]
  #    on_update   = CASCADE
  #    on_delete   = CASCADE
  #  }
#  foreign_key "solar_system_id" {
#    columns     = [column.solar_system_id]
#    ref_columns = [table.universe_system.column.id]
#    on_update   = CASCADE
#    on_delete   = CASCADE
#  }
}

table "universe_graphics" {
  schema = schema.public
  column "id" {
    type = int
  }
  column "collision_file" {
    type = text
    null = true
  }
  column "graphic_file" {
    type = text
    null = true
  }
  column "icon_folder" {
    type = text
    null = true
  }
  column "sof_dna" {
    type = text
    null = true
  }
  column "sof_fation_name" {
    type = text
    null = true
  }
  column "sof_hull_name" {
    type = text
    null = true
  }
  column "sof_race_name" {
    type = text
    null = true
  }
}

table "universe_group" {
  schema = schema.public
  column "id" {
    type = int
  }
  column "category_id" {
    type = int
  }
  column "name" {
    type = text
  }
  column "published" {
    type = bool
  }
  primary_key {
    columns = [column.id]
  }
#  foreign_key "category_id" {
#    columns     = [column.category_id]
#    ref_columns = [table.universe_category.column.id]
#    on_update   = CASCADE
#    on_delete   = CASCADE
#  }
}

table "universe_race" {
  schema = schema.public
  column "id" {
    type = int
  }
  column "alliance_id" {
    type = int
  }
  column "description" {
    type = text
  }
  column "name" {
    type = text
  }
  primary_key {
    columns = [column.id]
  }
}

table "universe_region" {
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
  primary_key {
    columns = [column.id]
  }
}

table "universe_stargate" {
  schema = schema.public
  column "id" {
    type = int
  }
  column "name" {
    type = text
  }
  column "destination_stargate_id" {
    type = int
  }
  column "destination_system_id" {
    type = int
  }
  column "position_x" {
    type = float
  }
  column "position_y" {
    type = float
  }
  column "position_z" {
    type = float
  }
  column "type_id" {
    type = int
  }
  column "system_id" {
    type = int
  }
  primary_key {
    columns = [column.id]
  }
#  foreign_key "destination_system_id" {
#    columns     = [column.destination_system_id]
#    ref_columns = [table.universe_system.column.id]
#    on_update   = CASCADE
#    on_delete   = CASCADE
#  }
#  foreign_key "type_id" {
#    columns     = [column.type_id]
#    ref_columns = [table.universe_type.column.id]
#    on_update   = CASCADE
#    on_delete   = CASCADE
#  }
}

table "universe_star" {
  schema = schema.public
  column "id" {
    type = int
  }
  column "name" {
    type = text
  }
  column "age" {
    type = bigint
  }
  column "luminosity" {
    type = float
  }
  column "radius" {
    type = bigint
  }
  column "solar_system_id" {
    type = int
  }
  column "spectral_class" {
    type = text
  }
  column "temperature" {
    type = int
  }
  column "type_id" {
    type = int
  }
  primary_key {
    columns = [column.id]
  }
#  foreign_key "solar_system_id" {
#    columns     = [column.solar_system_id]
#    ref_columns = [table.universe_system.column.id]
#    on_update   = CASCADE
#    on_delete   = CASCADE
#  }
#  foreign_key "type_id" {
#    columns     = [column.type_id]
#    ref_columns = [table.universe_type.column.id]
#    on_update   = CASCADE
#    on_delete   = CASCADE
#  }
}

table "universe_station" {
  schema = schema.public
  column "id" {
    type = int
  }
  column "name" {
    type = text
  }
  column "owner" {
    type = int
  }
  column "max_dockable_ship_volume" {
    type = float
  }
  column "office_rental_cost" {
    type = float
  }
  column "position_x" {
    type = float
  }
  column "position_y" {
    type = float
  }
  column "position_z" {
    type = float
  }
  column "race_id" {
    type = int
  }
  column "reprocessing_efficiency" {
    type = float
  }
  column "reprocessing_stations_take" {
    type = float
  }
  column "system_id" {
    type = int
  }
  column "type_id" {
    type = int
  }
  primary_key {
    columns = [column.id]
  }
#  foreign_key "system_id" {
#    columns     = [column.system_id]
#    ref_columns = [table.universe_system.column.id]
#    on_update   = CASCADE
#    on_delete   = CASCADE
#  }
#  foreign_key "type_id" {
#    columns     = [column.type_id]
#    ref_columns = [table.universe_type.column.id]
#    on_update   = CASCADE
#    on_delete   = CASCADE
#  }
}

table "universe_station_service" {
  schema = schema.public
  column "station_id" {
    type = int
  }
  column "service" {
    type = text
  }
  primary_key {
    columns = [column.station_id, column.service]
  }
  foreign_key "station_id" {
    columns     = [column.station_id]
    ref_columns = [table.universe_station.column.id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
}

table "universe_structure" {
  schema = schema.public
  column "id" {
    type = bigint
  }
  column "name" {
    type = text
  }
  column "owner_id" {
    type = int
  }
  column "position_x" {
    type = float
  }
  column "position_y" {
    type = float
  }
  column "position_z" {
    type = float
  }
  column "solar_system_id" {
    type = int
  }
  column "type_id" {
    type = int
  }
  primary_key {
    columns = [column.id]
  }
#  foreign_key "solar_system_id" {
#    columns     = [column.solar_system_id]
#    ref_columns = [table.universe_system.column.id]
#    on_update   = CASCADE
#    on_delete   = CASCADE
#  }
#  foreign_key "type_id" {
#    columns     = [column.type_id]
#    ref_columns = [table.universe_type.column.id]
#    on_update   = CASCADE
#    on_delete   = CASCADE
#  }
}

table "universe_system_jump" {
  schema = schema.public
  column "system_id" {
    type = int
  }
  column "timestamp" {
    type    = timestamptz
    default = sql("now()")
  }
  column "ship_jumps" {
    type = int
  }
  primary_key {
    columns = [column.system_id, column.timestamp]
  }
  foreign_key "system_id" {
    columns     = [column.system_id]
    ref_columns = [table.universe_system.column.id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
}

table "universe_system_kill" {
  schema = schema.public
  column "system_id" {
    type = int
  }
  column "timestamp" {
    type    = timestamptz
    default = sql("now()")
  }
  column "ship_kills" {
    type = int
  }
  column "pod_kills" {
    type = int
  }
  column "npc_kills" {
    type = int
  }
  primary_key {
    columns = [column.system_id, column.timestamp]
  }
  foreign_key "system_id" {
    columns     = [column.system_id]
    ref_columns = [table.universe_system.column.id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
}

table "market_history" {
  schema = schema.public
  column "type_id" {
    type = int
  }
  column "region_id" {
    type = int
  }
  column "date" {
    type = date
  }
  column "average" {
    type = float
  }
  column "highest" {
    type = float
  }
  column "lowest" {
    type = float
  }
  column "order_count" {
    type = int
  }
  column "volume" {
    type = bigint
  }
  primary_key {
    columns = [column.type_id, column.region_id, column.date]
  }
  foreign_key "type_id" {
    columns     = [column.type_id]
    ref_columns = [table.universe_type.column.id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
  foreign_key "region_id" {
    columns     = [column.region_id]
    ref_columns = [table.universe_region.column.id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
}

table "market_orders" {
  schema = schema.public
  column "id" {
    type = bigint
  }
  column "type_id" {
    type = int
  }
  column "region_id" {
    type = int
  }
  column "location_id" {
    type = bigint
  }
  column "range" {
    type = text
  }
  column "is_buy_order" {
    type = bool
  }
  column "price" {
    type = numeric
  }
  column "volume_total" {
    type = int
  }
  column "volume_remain" {
    type = int
  }
  column "issued" {
    type = timestamptz
  }
  column "duration" {
    type = int
  }
  column "min_volume" {
    type = int
  }
  column "system_id" {
    type = int
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "type_id" {
    columns     = [column.type_id]
    ref_columns = [table.universe_type.column.id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
  foreign_key "region_id" {
    columns     = [column.region_id]
    ref_columns = [table.universe_region.column.id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
  foreign_key "system_id" {
    columns     = [column.system_id]
    ref_columns = [table.universe_system.column.id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
}

table "market_price" {
  schema = schema.public
  column "type_id" {
    type = int
  }
  column "average_price" {
    type = numeric
    null = true
  }
  column "adjusted_price" {
    type = numeric
    null = true
  }
  column "timestamp" {
    type    = timestamptz
    default = sql("now()")
  }
  primary_key {
    columns = [column.type_id, column.timestamp]
  }
  foreign_key "type_id" {
    columns     = [column.type_id]
    ref_columns = [table.universe_type.column.id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
}

table "market_group" {
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
  column "parent_group_id" {
    type = int
    null = true
  }
  primary_key {
    columns = [column.id]
  }
}

table "alliance" {
  schema = schema.public
  column "id" {
    type = int
  }
  column "creator_corporation_id" {
    type = int
  }
  column "creator_id" {
    type = int
  }
  column "date_founded" {
    type = timestamptz
  }
  column "executor_corporation_id" {
    type = int
    null = true
  }
  column "faction_id" {
    type = int
    null = true
  }
  column "name" {
    type = text
  }
  column "ticker" {
    type = text
  }
  primary_key {
    columns = [column.id]
  }
  #  foreign_key "creator_corporation_id" {
  #    columns     = [column.creator_corporation_id]
  #    ref_columns = [table.corporation.column.id]
  #    on_update   = CASCADE
  #    on_delete   = CASCADE
  #  }
  #  foreign_key "creator_id" {
  #    columns     = [column.creator_id]
  #    ref_columns = [table.character.column.id]
  #    on_update   = CASCADE
  #    on_delete   = CASCADE
  #  }
  #  foreign_key "executor_corporation_id" {
  #    columns     = [column.executor_corporation_id]
  #    ref_columns = [table.corporation.column.id]
  #    on_update   = CASCADE
  #    on_delete   = CASCADE
  #  }
  #  foreign_key "faction_id" {
  #    columns     = [column.faction_id]
  #    ref_columns = [table.universe_faction.column.id]
  #    on_update   = CASCADE
  #    on_delete   = CASCADE
  #  }
}

table "alliance_icon" {
  schema = schema.public
  column "id" {
    type = int
  }
  column "px128x128" {
    type = text
  }
  column "px64x64" {
    type = text
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "alliance_id" {
    columns     = [column.id]
    ref_columns = [table.alliance.column.id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
}

table "corporation" {
  schema = schema.public
  column "id" {
    type = int
  }
  column "alliance_id" {
    type = int
    null = true
  }
  column "ceo_id" {
    type = int
  }
  column "creator_id" {
    type = int
  }
  column "date_founded" {
    type = timestamptz
  }
  column "description" {
    type = text
  }
  column "faction_id" {
    type = int
    null = true
  }
  column "home_station_id" {
    type = int
    null = true
  }
  column "member_count" {
    type = int
  }
  column "name" {
    type = text
  }
  column "shares" {
    type = int
  }
  column "tax_rate" {
    type = float
  }
  column "ticker" {
    type = text
  }
  column "url" {
    type = text
    null = true
  }
  column "war_eligible" {
    type    = bool
    default = false
  }
  primary_key {
    columns = [column.id]
  }
  #  foreign_key "alliance_id" {
  #    columns     = [column.alliance_id]
  #    ref_columns = [table.alliance.column.id]
  #    on_update   = CASCADE
  #    on_delete   = CASCADE
  #  }
  #  foreign_key "ceo_id" {
  #    columns     = [column.ceo_id]
  #    ref_columns = [table.character.column.id]
  #    on_update   = CASCADE
  #    on_delete   = CASCADE
  #  }
  #  foreign_key "creator_id" {
  #    columns     = [column.creator_id]
  #    ref_columns = [table.character.column.id]
  #    on_update   = CASCADE
  #    on_delete   = CASCADE
  #  }
  #  foreign_key "faction_id" {
  #    columns     = [column.faction_id]
  #    ref_columns = [table.universe_faction.column.id]
  #    on_update   = CASCADE
  #    on_delete   = CASCADE
  #  }
  #  foreign_key "home_station_id" {
  #    columns     = [column.home_station_id]
  #    ref_columns = [table.universe_station.column.id]
  #    on_update   = CASCADE
  #    on_delete   = CASCADE
  #  }
}

table "corporation_alliance_history" {
  schema = schema.public
  column "id" {
    type = int
  }
  column "alliance_id" {
    type = int
  }
  column "is_deleted" {
    type    = bool
    default = false
  }
  column "record_id" {
    type = int
  }
  column "start_date" {
    type = timestamptz
  }
  primary_key {
    columns = [column.id, column.record_id]
  }
  #  foreign_key "alliance_id" {
  #    columns     = [column.alliance_id]
  #    ref_columns = [table.alliance.column.id]
  #    on_update   = CASCADE
  #    on_delete   = CASCADE
  #  }
  foreign_key "corporation_id" {
    columns     = [column.id]
    ref_columns = [table.corporation.column.id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
}

table "corporation_icon" {
  schema = schema.public
  column "id" {
    type = int
  }
  column "px128x128" {
    type = text
    null = true
  }
  column "px256x256" {
    type = text
    null = true
  }
  column "px64x64" {
    type = text
    null = true
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "corporation_id" {
    columns     = [column.id]
    ref_columns = [table.corporation.column.id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
}

table "character" {
  schema = schema.public
  column "id" {
    type = int
  }
  column "alliance_id" {
    type = int
    null = true
  }
  column "birthday" {
    type = timestamptz
  }
  column "bloodline_id" {
    type = int
  }
  column "corporation_id" {
    type = int
  }
  column "description" {
    type = text
    null = true
  }
  column "faction_id" {
    type = int
    null = true
  }
  column "gender" {
    type = text
  }
  column "name" {
    type = text
  }
  column "race_id" {
    type = int
  }
  column "security_status" {
    type = float
    null = true
  }
  column "title" {
    type = text
    null = true
  }
  primary_key {
    columns = [column.id]
  }
  #  foreign_key "alliance_id" {
  #    columns     = [column.alliance_id]
  #    ref_columns = [table.alliance.column.id]
  #    on_update   = CASCADE
  #    on_delete   = CASCADE
  #  }
  #  foreign_key "bloodline_id" {
  #    columns     = [column.bloodline_id]
  #    ref_columns = [table.universe_bloodlines.column.id]
  #    on_update   = CASCADE
  #    on_delete   = CASCADE
  #  }
  #  foreign_key "corporation_id" {
  #    columns     = [column.corporation_id]
  #    ref_columns = [table.corporation.column.id]
  #    on_update   = CASCADE
  #    on_delete   = CASCADE
  #  }
  #  foreign_key "faction_id" {
  #    columns     = [column.faction_id]
  #    ref_columns = [table.universe_faction.column.id]
  #    on_update   = CASCADE
  #    on_delete   = CASCADE
  #  }
}

table "character_corporation_history" {
  schema = schema.public
  column "id" {
    type = int
  }
  column "corporation_id" {
    type = int
  }
  column "is_deleted" {
    type    = bool
    default = false
  }
  column "record_id" {
    type = int
  }
  column "start_date" {
    type = timestamptz
  }
  primary_key {
    columns = [column.id, column.record_id]
  }
  #  foreign_key "corporation_id" {
  #    columns     = [column.corporation_id]
  #    ref_columns = [table.corporation.column.id]
  #    on_update   = CASCADE
  #    on_delete   = CASCADE
  #  }
  foreign_key "character_id" {
    columns     = [column.id]
    ref_columns = [table.character.column.id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
}

table "character_portrait" {
  schema = schema.public
  column "id" {
    type = int
  }
  column "px128x128" {
    type = text
  }
  column "px256x256" {
    type = text
  }
  column "px512x512" {
    type = text
  }
  column "px64x64" {
    type = text
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "character_id" {
    columns     = [column.id]
    ref_columns = [table.character.column.id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
}

table "contract" {
  schema = schema.public
  column "id" {
    type = bigint
  }
  column "buyout" {
    type = numeric
    null = true
  }
  column "collateral" {
    type = numeric
    null = true
  }
  column "date_completed" {
    type = timestamptz
    null = true
  }
  column "date_expired" {
    type = timestamptz
  }
  column "date_issued" {
    type = timestamptz
  }
  column "days_to_complete" {
    type = int
    null = true
  }
  column "end_location_id" {
    type = bigint
    null = true
  }
  column "for_corporation" {
    type    = bool
    default = false
  }
  column "issuer_corporation_id" {
    type = int
  }
  column "issuer_id" {
    type = int
  }
  column "price" {
    type = numeric
    null = true
  }
  column "reward" {
    type = numeric
    null = true
  }
  column "start_location_id" {
    type = bigint
    null = true
  }
  column "status" {
    type = text
  }
  column "title" {
    type = text
    null = true
  }
  column "type" {
    type = text
  }
  column "volume" {
    type = numeric
    null = true
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "issuer_corporation_id" {
    columns     = [column.issuer_corporation_id]
    ref_columns = [table.corporation.column.id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
  foreign_key "issuer_id" {
    columns     = [column.issuer_id]
    ref_columns = [table.character.column.id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
  #  foreign_key "acceptor_id" {
  #    columns     = [column.acceptor_id]
  #    ref_columns = [table.character.column.id]
  #    on_update   = CASCADE
  #    on_delete   = CASCADE
  #  }
  #  foreign_key "assignee_id" {
  #    columns     = [column.assignee_id]
  #    ref_columns = [table.character.column.id]
  #    on_update   = CASCADE
}

table "contract_bid" {
  schema = schema.public
  column "id" {
    type = bigint
  }
  column "amount" {
    type = numeric
  }
  column "bid_id" {
    type = int
  }
  column "date_bid" {
    type = timestamptz
  }
  primary_key {
    columns = [column.id, column.bid_id]
  }
  foreign_key "contract_id" {
    columns     = [column.id]
    ref_columns = [table.contract.column.id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
}

table "contract_item" {
  schema = schema.public
  column "id" {
    type = int
  }
  column "is_included" {
    type = bool
  }
  column "is_blueprint_copy" {
    type = bool
    null = true
  }
  column "item_id" {
    type = bigint
  }
  column "material_efficiency" {
    type = int
    null = true
  }
  column "quantity" {
    type = int
  }
  column "record_id" {
    type = int
  }
  column "runs" {
    type = int
    null = true
  }
  column "time_efficiency" {
    type = int
    null = true
  }
  primary_key {
    columns = [column.id, column.record_id]
  }
  foreign_key "contract_id" {
    columns     = [column.id]
    ref_columns = [table.contract.column.id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
  foreign_key "item_id" {
    columns     = [column.item_id]
    ref_columns = [table.universe_type.column.id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
}
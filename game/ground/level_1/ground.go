components {
  id: "ground"
  component: "/game/ground/level_1/ground.tilemap"
}
components {
  id: "ground_viz"
  component: "/game/ground/level_1/ground_viz.tilemap"
  position {
    z: 0.1
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "collision_shape: \"/game/ground/level_1/ground.tilemap\"\n"
  "type: COLLISION_OBJECT_TYPE_STATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"default\"\n"
  "mask: \"default\"\n"
  "mask: \"player\"\n"
  "mask: \"player_bullet\"\n"
  "mask: \"enemy\"\n"
  "mask: \"enemy_bullet\"\n"
  ""
}

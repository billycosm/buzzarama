embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"bullet_player_64_64\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/bullet_player.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 0.2
    y: 0.2
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_DYNAMIC\n"
  "mass: 0.05\n"
  "friction: 0.0\n"
  "restitution: 1.0\n"
  "group: \"player_bullet\"\n"
  "mask: \"enemy\"\n"
  "mask: \"default\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 1\n"
  "    id: \"bullet_collider\"\n"
  "  }\n"
  "  data: 4.442382\n"
  "}\n"
  "locked_rotation: true\n"
  "bullet: true\n"
  ""
}

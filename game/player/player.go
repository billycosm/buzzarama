components {
  id: "player_control"
  component: "/game/player/player_control.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"player_64_64\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/player.atlas\"\n"
  "}\n"
  ""
}

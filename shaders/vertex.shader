#version 460

layout (location = 0) in uvec3 position;
layout (location = 1) in uint  color;

out vec3 out_color;

void main () {
	gl_Position = vec4(vertex + offset, 1.0);
	out_color = color;
}

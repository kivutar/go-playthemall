package video

var fontVertexShader = `
#if __VERSION__ >= 130
#define COMPAT_VARYING out
#define COMPAT_ATTRIBUTE in
#define COMPAT_TEXTURE texture
#else
#define COMPAT_VARYING varying
#define COMPAT_ATTRIBUTE attribute
#define COMPAT_TEXTURE texture2D
#endif

COMPAT_ATTRIBUTE vec2 vert;
COMPAT_ATTRIBUTE vec2 vertTexCoord;
uniform vec2 resolution;
COMPAT_VARYING vec2 fragTexCoord;

void main() {
	// convert the rectangle from pixels to 0.0 to 1.0
	vec2 zeroToOne = vert / resolution;
	// convert from 0->1 to 0->2
	vec2 zeroToTwo = zeroToOne * 2.0;
	// convert from 0->2 to -1->+1 (clipspace)
	vec2 clipSpace = zeroToTwo - 1.0;
	fragTexCoord = vertTexCoord;
	gl_Position = vec4(clipSpace * vec2(1, -1), 0, 1);
}
` + "\x00"

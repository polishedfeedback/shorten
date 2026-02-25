const manifest = (() => {
function __memo(fn) {
	let value;
	return () => value ??= (value = fn());
}

return {
	appDir: "_app",
	appPath: "_app",
	assets: new Set(["robots.txt"]),
	mimeTypes: {".txt":"text/plain"},
	_: {
		client: {start:"_app/immutable/entry/start.CUIu-_xh.js",app:"_app/immutable/entry/app.Cjbtob1H.js",imports:["_app/immutable/entry/start.CUIu-_xh.js","_app/immutable/chunks/7lrU0nzN.js","_app/immutable/chunks/BQNkOFaZ.js","_app/immutable/chunks/BllHz6Nw.js","_app/immutable/entry/app.Cjbtob1H.js","_app/immutable/chunks/BQNkOFaZ.js","_app/immutable/chunks/DXC2exyq.js","_app/immutable/chunks/CgIzCHWt.js","_app/immutable/chunks/BllHz6Nw.js","_app/immutable/chunks/lq4UHFV3.js","_app/immutable/chunks/DNcYQha-.js"],stylesheets:[],fonts:[],uses_env_dynamic_public:false},
		nodes: [
			__memo(() => import('./chunks/0-C97TqbwD.js')),
			__memo(() => import('./chunks/1-BiGpUz8n.js')),
			__memo(() => import('./chunks/2-CwF4_gh5.js')),
			__memo(() => import('./chunks/3-B0YudSHW.js'))
		],
		remotes: {
			
		},
		routes: [
			{
				id: "/",
				pattern: /^\/$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 2 },
				endpoint: null
			},
			{
				id: "/dashboard",
				pattern: /^\/dashboard\/?$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 3 },
				endpoint: null
			}
		],
		prerendered_routes: new Set([]),
		matchers: async () => {
			
			return {  };
		},
		server_assets: {}
	}
}
})();

const prerendered = new Set([]);

const base = "";

export { base, manifest, prerendered };
//# sourceMappingURL=manifest.js.map

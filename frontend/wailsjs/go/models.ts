export namespace dto {
	
	export class NewConnectionReq {
	    name: string;
	    urls: string;
	    username: string;
	    password: string;
	
	    static createFrom(source: any = {}) {
	        return new NewConnectionReq(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.urls = source["urls"];
	        this.username = source["username"];
	        this.password = source["password"];
	    }
	}

}


export namespace models {
	
	export class BaseResponse {
	    err_code: string;
	    err_msg: string;
	    data: any;
	
	    static createFrom(source: any = {}) {
	        return new BaseResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.err_code = source["err_code"];
	        this.err_msg = source["err_msg"];
	        this.data = source["data"];
	    }
	}
	export class CatIndexReq {
	    id: string;
	
	    static createFrom(source: any = {}) {
	        return new CatIndexReq(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	    }
	}
	
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
	export class QueryReq {
	    page_number: number;
	    page_size: number;
	    connection_id: string;
	    index: string;
	    condition: {[key: string]: any};
	
	    static createFrom(source: any = {}) {
	        return new QueryReq(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.page_number = source["page_number"];
	        this.page_size = source["page_size"];
	        this.connection_id = source["connection_id"];
	        this.index = source["index"];
	        this.condition = source["condition"];
	    }
	}

}


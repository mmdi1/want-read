export namespace db {
	
	export class Book {
	    id: string;
	    name: string;
	    file_name: string;
	    // Go type: time.Time
	    created_at: any;
	    total_size: number;
	    read_size: number;
	
	    static createFrom(source: any = {}) {
	        return new Book(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.file_name = source["file_name"];
	        this.created_at = this.convertValues(source["created_at"], null);
	        this.total_size = source["total_size"];
	        this.read_size = source["read_size"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}


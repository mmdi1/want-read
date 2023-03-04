export namespace db {
	
	export class Book {
	    id: string;
	    name: string;
	    file_name: string;
	    // Go type: time.Time
	    update_at: any;
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
	        this.update_at = this.convertValues(source["update_at"], null);
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
	export class Setting {
	    prev_group: number[];
	    next_group: number[];
	    hide_group: number[];
	    font_size: number;
	    font_color: string;
	    show_size: number;
	    read_width: number;
	    read_height: number;
	
	    static createFrom(source: any = {}) {
	        return new Setting(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.prev_group = source["prev_group"];
	        this.next_group = source["next_group"];
	        this.hide_group = source["hide_group"];
	        this.font_size = source["font_size"];
	        this.font_color = source["font_color"];
	        this.show_size = source["show_size"];
	        this.read_width = source["read_width"];
	        this.read_height = source["read_height"];
	    }
	}

}


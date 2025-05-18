export namespace main {
	
	export class OpenedFile {
	    IsModified: boolean;
	    FileContent: string;
	    Err: boolean;
	
	    static createFrom(source: any = {}) {
	        return new OpenedFile(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.IsModified = source["IsModified"];
	        this.FileContent = source["FileContent"];
	        this.Err = source["Err"];
	    }
	}

}


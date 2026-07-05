export namespace backend {
	
	export class LPreference {
	    LPreferenceInput: string[];
	    LPreferenceOutput: string;
	    LPreferenceMirror: boolean;
	    LPreferenceTree: boolean;
	    LPreferenceSuffix: string;
	    LPreferenceCaution: boolean;
	    LPreferenceWarning: boolean;
	    LPreferenceMarker: string;
	    LPreferencePattern: string;
	    LPreferenceCustom: boolean;
	    LPreferenceUnnumbered: boolean;
	    LPreferenceFFmpeg: string;
	    LPreferenceTemporary: string;
	
	    static createFrom(source: any = {}) {
	        return new LPreference(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.LPreferenceInput = source["LPreferenceInput"];
	        this.LPreferenceOutput = source["LPreferenceOutput"];
	        this.LPreferenceMirror = source["LPreferenceMirror"];
	        this.LPreferenceTree = source["LPreferenceTree"];
	        this.LPreferenceSuffix = source["LPreferenceSuffix"];
	        this.LPreferenceCaution = source["LPreferenceCaution"];
	        this.LPreferenceWarning = source["LPreferenceWarning"];
	        this.LPreferenceMarker = source["LPreferenceMarker"];
	        this.LPreferencePattern = source["LPreferencePattern"];
	        this.LPreferenceCustom = source["LPreferenceCustom"];
	        this.LPreferenceUnnumbered = source["LPreferenceUnnumbered"];
	        this.LPreferenceFFmpeg = source["LPreferenceFFmpeg"];
	        this.LPreferenceTemporary = source["LPreferenceTemporary"];
	    }
	}
	export class LReportMetric {
	    LReportLabel: string;
	    LReportValue: string;
	
	    static createFrom(source: any = {}) {
	        return new LReportMetric(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.LReportLabel = source["LReportLabel"];
	        this.LReportValue = source["LReportValue"];
	    }
	}
	export class LReportSection {
	    LReportTitle: string;
	    LReportTag: string;
	    LReportBadge: string;
	    LReportItem: string[];
	    LReportMetric: LReportMetric[];
	
	    static createFrom(source: any = {}) {
	        return new LReportSection(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.LReportTitle = source["LReportTitle"];
	        this.LReportTag = source["LReportTag"];
	        this.LReportBadge = source["LReportBadge"];
	        this.LReportItem = source["LReportItem"];
	        this.LReportMetric = this.convertValues(source["LReportMetric"], LReportMetric);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class LReportFile {
	    LReportNumber: number;
	    LReportName: string;
	    LReportPath: string;
	    LReportAsset: string;
	    LReportDurationSecond: number;
	
	    static createFrom(source: any = {}) {
	        return new LReportFile(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.LReportNumber = source["LReportNumber"];
	        this.LReportName = source["LReportName"];
	        this.LReportPath = source["LReportPath"];
	        this.LReportAsset = source["LReportAsset"];
	        this.LReportDurationSecond = source["LReportDurationSecond"];
	    }
	}
	export class LReportGroup {
	    LReportKey: string;
	    LReportName: string;
	    LReportDirectory: string;
	    LReportSize: string;
	    LReportDuration: string;
	    LReportLoudness: string;
	    LReportCompatibility: string;
	    LReportCompatibilityTag: string;
	    LReportTask: string;
	    LReportTaskTag: string;
	    LReportOutputTitle: string;
	    LReportOutputText: string;
	    LReportFile: LReportFile[];
	    LReportSection: LReportSection[];
	
	    static createFrom(source: any = {}) {
	        return new LReportGroup(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.LReportKey = source["LReportKey"];
	        this.LReportName = source["LReportName"];
	        this.LReportDirectory = source["LReportDirectory"];
	        this.LReportSize = source["LReportSize"];
	        this.LReportDuration = source["LReportDuration"];
	        this.LReportLoudness = source["LReportLoudness"];
	        this.LReportCompatibility = source["LReportCompatibility"];
	        this.LReportCompatibilityTag = source["LReportCompatibilityTag"];
	        this.LReportTask = source["LReportTask"];
	        this.LReportTaskTag = source["LReportTaskTag"];
	        this.LReportOutputTitle = source["LReportOutputTitle"];
	        this.LReportOutputText = source["LReportOutputText"];
	        this.LReportFile = this.convertValues(source["LReportFile"], LReportFile);
	        this.LReportSection = this.convertValues(source["LReportSection"], LReportSection);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class LReport {
	    LReportKind: string;
	    LTaskMessage: string;
	    LTaskCancel: boolean;
	    LProgressTotal: number;
	    LProgressProcessed: number;
	    LProgressPercent: number;
	    LReportGroup: LReportGroup[];
	
	    static createFrom(source: any = {}) {
	        return new LReport(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.LReportKind = source["LReportKind"];
	        this.LTaskMessage = source["LTaskMessage"];
	        this.LTaskCancel = source["LTaskCancel"];
	        this.LProgressTotal = source["LProgressTotal"];
	        this.LProgressProcessed = source["LProgressProcessed"];
	        this.LProgressPercent = source["LProgressPercent"];
	        this.LReportGroup = this.convertValues(source["LReportGroup"], LReportGroup);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	
	
	
	
	export class LTemporaryResult {
	    LTemporaryPath: string;
	    LTemporaryCount: number;
	
	    static createFrom(source: any = {}) {
	        return new LTemporaryResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.LTemporaryPath = source["LTemporaryPath"];
	        this.LTemporaryCount = source["LTemporaryCount"];
	    }
	}

}

export namespace bridge {
	
	export class LProgramProfile {
	    LProgramName: string;
	    LProgramVersion: string;
	    LProgramAuthorName: string;
	    LProgramAuthorEmail: string;
	
	    static createFrom(source: any = {}) {
	        return new LProgramProfile(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.LProgramName = source["LProgramName"];
	        this.LProgramVersion = source["LProgramVersion"];
	        this.LProgramAuthorName = source["LProgramAuthorName"];
	        this.LProgramAuthorEmail = source["LProgramAuthorEmail"];
	    }
	}

}


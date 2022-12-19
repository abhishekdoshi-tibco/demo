
import {Injectable, Injector} from "@angular/core";
import {Http} from "@angular/http";
import {Observable} from "rxjs/Observable";
import {
    WiContrib,
    WiServiceHandlerContribution,
    IValidationResult,
    ValidationResult,
    IContributionTypes,
    ActionResult,
    IActionResult,
    WiContribModelService,
    WiContributionUtils,
    IConnectorContribution
} from "wi-studio/app/contrib/wi-contrib";

@WiContrib({})
@Injectable()
export class createHandler extends WiServiceHandlerContribution {

    constructor(private injector: Injector, private http: Http) {
                    super(injector, http);
                }

    value = (fieldName: string, context: IContributionTypes): Observable<any> | any => {
        return Observable.create(observer => {
            observer.next("");
        });
    }

    validate = (fieldName: string, context: IContributionTypes): Observable<IValidationResult> | IValidationResult => {
        return Observable.create(observer => {
            let vresult: IValidationResult = ValidationResult.newValidationResult();
            observer.next(vresult);
        });
    }

    action = (actionId: string, context: IContributionTypes): Observable<IActionResult> | IActionResult => {
        return Observable.create(observer => {
            let aresult: IActionResult = ActionResult.newActionResult();
            observer.next(aresult);
        });
    }
}


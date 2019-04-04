import { iocContainer, ILogger, Logger } from "@pyramidlabs/core";
import { UserController } from "./user-controller";
import { BudgetRepository } from "@pyramidlabs/components";
import { TestController } from "./test-controller";
import { AuthenticationController } from "./authentication-controller";
import { interfaces } from "inversify";

export class ContainerOrchestrator {
    static initializeContainer = () => {
        iocContainer.bind(BudgetRepository).to(BudgetRepository);
        iocContainer.bind(TestController).to(TestController);
        iocContainer.bind(AuthenticationController).to(AuthenticationController);
        iocContainer
            .bind(ILogger)
            .toDynamicValue((context: interfaces.Context) => {
                return new Logger("{{.serviceName}}"); // This will be in the template at least
            });
        console.log(iocContainer);
    };
}

import { Controller } from "tsoa";
import {
    Container,
    decorate,
    injectable as Injectable,
    inject as Inject,
    interfaces
} from "inversify";
import {
    autoProvide as AutoProvide,
    fluentProvide,
    buildProviderModule,
    provide as Provide

} from "inversify-binding-decorators";
import 'reflect-metadata';

decorate(Injectable(), Controller);

const iocContainer = new Container();

type Identifier =
    | string
    | symbol
    | interfaces.Newable<any>
    | interfaces.Abstract<any>;

const ProvideNamed = (identifier: Identifier, name: string) => {
    return fluentProvide(identifier)
        .whenTargetNamed(name)
        .done();
};

const ProvideSingleton = (identifier: Identifier) => {
    return fluentProvide(identifier)
        .inSingletonScope()
        .done(true);
};

// iocContainer.load(buildProviderModule());
console.log('ioc');
export {
    iocContainer,
    AutoProvide,
    Provide,
    ProvideSingleton,
    ProvideNamed,
    Inject,
    Injectable
};

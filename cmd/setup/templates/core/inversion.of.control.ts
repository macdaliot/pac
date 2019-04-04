import { Controller } from "tsoa";
import {
    decorate,
    Container,
    injectable as Injectable,
    inject as Inject
} from "inversify";
import "reflect-metadata";


decorate(Injectable(), Controller);

const iocContainer = new Container();

export {
    iocContainer,
    Inject,
    Injectable
};

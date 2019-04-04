import * as express from "express";
import {
    Route,
    Get,
    Post,
    Controller,
    Request,
    Query
} from "tsoa";
import { Injectable } from "@pyramidlabs/core";
import { createJwt } from "../../core/passport-strategies";
console.log("authentication");

@Injectable()
@Route("auth")
export class AuthenticationController extends Controller {
    @Get("login")
    get(@Query() jwtToken: string) {
        return jwtToken;
    }

    @Post("callback")
    login(@Request() request: express.Request) {
        const response = request.res;
        createJwt(request, (err, encoded) => {
            if (err) {
                response.status(500).send(err);
            } else {
                response.redirect(process.env.APP_ROOT + "/api/auth/login?" + encoded);
            }
        });
    }
}

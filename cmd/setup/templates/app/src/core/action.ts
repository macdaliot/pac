/**
 * Use this type definition instead of `Function` type constructor
 */
export type AnyFunction = (...args: any[]) => any;

/**
 * Simple alias to save you keystrokes when defining JS typed object maps
 */
export type StringMap<T> = { [key: string]: T };

export type Action<T extends string, P = void> = P extends void
    ? Readonly<{ type: T }>
    : Readonly<{ type: T; payload: P }>;



/**
 * Get actions types union when using the create action function.
 * It is recommended to export them under the same name as your Actions object, to leverage token merging
 *
 * @example
 *
 * ```ts
 * const GET_USER_ID = 'GET_USER_ID'
 * const SET_USER = 'SET_USER'
 *
 * export const Actions = {
 *  setAge: (age: number) => createAction(GET_USER, age),
 *  setName: (name: string) => createAction(SET_USER, name),
 *  reloadUrl: () => createAction(RELOAD_URL),
 * }
 *
 * // The type will be the following:
 * { Readonly <{ type: "GET_USER_ID"; payload: number; }> | Readonly <{ type: "SET_USER"; payload: string; }>
 * export type Actions = ActionsUnion<typeof Actions>
 * ```
 */
export type ActionsUnion<A extends StringMap<AnyFunction>> = ReturnType<
    A[keyof A]
    >;

/**
 * Only use if you prefer to use if else than switch statements
 * @param actionName The action name
 * @param action The action object
 */
export const isActionOf = (
    actionName: string,
    action: Readonly<{ type: string }>
): boolean => {
  return action.type === actionName;
};

export function createAction<T extends string>(type: T): Action<T>
export function createAction<T extends string, P>(
    type: T,
    payload: P
): Action<T, P>
export function createAction<T extends string, P>(type: T, payload?: P) {
  const action = payload === undefined ? { type } : { type, payload }

  return action
}

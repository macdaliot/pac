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
  : Readonly<{ type: string; payload: P }>;

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

export function createAction<T extends string>(type: T): Action<T>;
export function createAction<T extends string, P>(
  type: string,
  payload: P
): Action<T, P>;
export function createAction<T, P>(type: T, payload?: P) {
  const action = payload === undefined ? { type } : { type, payload };
  return action;
}

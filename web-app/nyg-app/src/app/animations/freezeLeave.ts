import { style, transition, trigger, query, group, animate, animateChild, keyframes } from "@angular/animations";
const timelineFunction = 'cubic-bezier(1, .71, 0, -0.06)'
const duration = '.45s'

export const onUnfreeze = trigger("OnUnfreeze", [
    transition("* <=> *", [
        group([

            query(
                ":leave",
                [
                    animate(
                        `${duration} ${timelineFunction}`,

                        keyframes([
                            style({ width: '0px', opacity: 0, offset: 1 })
                        ]),
                    ),
                ], { optional: true }
            ),
        ]),
    ]),
]);
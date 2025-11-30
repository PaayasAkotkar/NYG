import { style, transition, trigger, query, group, animate, animateChild, keyframes } from "@angular/animations";
const timelineFunction = 'cubic-bezier(1, .71, 0, -0.06)'
const duration = '.45s'

export const OnEnter = trigger("OnEnter", [
    transition("* <=> *", [
        group([
            query(
                ":enter",
                [

                    animate(
                        `${duration} ${timelineFunction}`,
                        keyframes([
                            style({ opacity: 0, transform: 'scale(1.2)', offset: 0 }),
                            style({ transform: 'scale(1)', offset: 1 }),
                        ])

                    ),
                ], { optional: true }
            ),
            query(
                ":leave",
                [
                    animate(
                        `${duration} ${timelineFunction}`,

                        keyframes([
                            style({ transform: 'translateY(0)', offset: 0 }),
                            style({ transform: 'translateY(-100%)', opacity: 0, offset: 1 })
                        ]),
                    ),
                ], { optional: true }
            ),
        ]),
    ]),
]);
import { animate, animation, keyframes, query, style, transition, trigger, group, stagger } from "@angular/animations";


const duration = '110ms '
const animationTimingFunction = 'cubic-bezier(.46,.31,.48,.14)'
const animationSpace = duration + animationTimingFunction

export const in_out = trigger('EnterLeave', [
    transition(':leave',
        [group([
            query('.pop-up', stagger('150ms', animate(animationSpace, keyframes([
                style({ transform: 'translate(0px,500vh)', offset: 0 }),
                style({ transform: 'translate(0px,600vh)', offset: 0.5 }),
                style({ visibilty: 'hidden', backgroundColor: 'red', offset: 1 })
            ])))),
        ])]),
])

export const myAnimation = trigger("myAnimation", [
    transition("* <=> *", [
        query(":enter, :leave", style({ position: "fixed", width: "100%" }), {
            optional: true,
        }),
        group([
            // block executes in parallel
            query(
                ":enter",
                [
                    style({ transform: "translateX(100%)" }),
                    animate("0.5s ease-in-out", style({ transform: "translateX(0%)" })),
                ],
                { optional: true }
            ),
            query(
                ":leave",
                [
                    style({ transform: "translateX(0%)" }),
                    animate(
                        "0.5s ease-in-out",
                        style({ transform: "translateX(-100%)" })
                    ),
                ],
                { optional: true }
            ),
        ]),
    ]),
]);
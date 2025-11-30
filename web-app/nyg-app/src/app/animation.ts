import { style, transition, trigger, query, group, animate, animateChild, keyframes } from "@angular/animations";
export const myAnimation = trigger("myAnimation", [
    transition("* <=> *", [
        group([
            query(
                ":enter",
                [

                    animate(
                        "0.945s cubic-bezier(.9, .81, .67, .68)",
                        keyframes([
                            style({ opacity: 0, offset: 0 }),
                            style({ transform: 'rotateX(30deg)', offset: 1 }),
                        ])

                    ),
                    animateChild(),
                ],
                { optional: true }
            ),
            query(
                ":leave",
                [
                    animate(".945s cubic-bezier(.9, .81, .67, .68)", keyframes([
                        style({ opacity: 0, offset: 1 }),
                    ]),
                    ),
                    animateChild()

                ],
                { optional: true }
            ),
        ]),
    ]),
]);
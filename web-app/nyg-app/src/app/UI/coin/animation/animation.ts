
import { group, trigger, query, animate, state, transition, style, keyframes } from '@angular/animations';



// export const FlipCoin = trigger(
//     'HeadTails', [
//     state('heads', style({ transform: 'rotateY(0deg)' })),
//     state('tails', style({ transform: 'rotateY(180deg)' })),
//     state('flip', style({})),

//     transition('heads => flip', [
//         animate('1s', keyframes([
//             style({ transform: 'rotateY(0deg)', offset: 0 }),       // Start at heads
//             style({ transform: 'rotateY(360deg)', offset: 0.5 }), // Halfway through (one full rotation)
//             style({ transform: 'rotateY(720deg)', offset: 1 }),
//         ])),
//     ]),

//     transition('flip=> heads', [
//         query('.nyg-head', [
//             animate('1s', keyframes([
//                 style({ transform: 'rotateY(720deg)' })
//             ]))
//         ]),

//     ]),
//     transition('flip=> tails', [
//         query('.nyg-tail', [
//             animate('1s', keyframes([
//                 style({ transform: 'rotateY(720deg)' })
//             ]))
//         ]),

//     ]),
// ]
// )

const timelineFunction = 'cubic-bezier(1, .71, 0, -0.06)'
const duration = '1.45s'

export const FlipCoin = trigger(
    'HeadTails', [
    state('heads', style({
    })),
    state('tails', style({
    })),
    transition("heads=>tails",
        group([
            query('.animate-tail', animate(`${duration} ${timelineFunction}`, keyframes([

                style({
                    transform: 'rotateX(180deg)', offset: 0
                }),

                style({
                    transform: 'rotateX(5turn)', offset: 1
                }),

            ]))),
            query('.animate-head', animate(`${duration} ${timelineFunction}`, keyframes([
                style({
                    transform: 'rotateX(0deg)', offset: 0
                }),

                style({
                    transform: 'rotateX(5turn)', offset: 1
                }),
            ]))),

        ]))
    ,
    transition("tails=>heads",
        group([
            query('.animate-tail', animate(`${duration} ${timelineFunction}`, keyframes([

                style({
                    transform: 'rotateX(180deg)', offset: 0
                }),

                style({
                    transform: 'rotateX(5turn)', offset: 1
                }),

            ]))),
            query('.animate-head', animate(`${duration} ${timelineFunction}`, keyframes([
                style({
                    transform: 'rotateX(0deg)', offset: 0
                }),

                style({
                    transform: 'rotateX(5turn)', offset: 1
                }),
            ]))),

        ]))
    ,

    // transition('heads=> tails',
    //     [
    //         animate('1s', keyframes([
    //             style({ transform: 'rotateX(180deg)', offset: 0 }),
    //             style({ transform: 'rotateX(11turn)', offset: 1 }),
    //         ])),
    //     ]),

    // transition('tails=> heads', [
    //     style({ transform: 'rotateX(180deg)', offset: 0 }),
    //     style({ transform: 'rotateX(11turn)', offset: 1 }),

    // ]),
    transition('*<=> *', [
        animate('1s',)

    ]),
]
)



export const X = trigger('childAnimation', [

    state(
        'true',
        style({
            width: '250px',
            opacity: 1,
            backgroundColor: 'yellow',
        }),
    ),
    state(
        'false',
        style({
            width: '100px',
            opacity: 0.8,
            backgroundColor: 'blue',
        }),
    ),
    transition('* => *', [animate('1s')]),
])
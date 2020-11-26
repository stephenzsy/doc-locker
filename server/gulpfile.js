'use strict';

const { task, series, src, dest } = require("gulp");
const ts = require("gulp-typescript");

task('build', () => {
    const tsProject = ts.createProject("tsconfig.json");

    return tsProject
        .src()
        .pipe(tsProject())
        .js
        .pipe(dest("dist"));
});

task('copyTemplates', () => {
    return src('templates/**/*.mustache')
        .pipe(dest('dist/templates'))
})

exports.default = series('build', 'copyTemplates');

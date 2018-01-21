#!/usr/bin/env python
# -*- coding: utf-8 -*-

import subprocess
import pytest


# Command line version, help, usage flag tests

def test_unilist_main_version_request_short():
    out = subprocess.check_output(
        "./unilist -v",
        stderr=subprocess.STDOUT,
        shell=True)

    assert "unilist v" in str(out)


def test_unilist_main_version_request_long():
    out = subprocess.check_output(
        "./unilist --version",
        stderr=subprocess.STDOUT,
        shell=True)

    assert "unilist v" in str(out)


def test_unilist_main_usage_request():
    out = subprocess.check_output(
        "./unilist --usage",
        stderr=subprocess.STDOUT,
        shell=True)

    assert "Usage:" in str(out)


def test_unilist_main_help_request_short():
    out = subprocess.check_output(
        "./unilist -h",
        stderr=subprocess.STDOUT,
        shell=True)

    assert "==========" in str(out)


def test_unilist_main_help_request_long():
    out = subprocess.check_output(
        "./unilist --help",
        stderr=subprocess.STDOUT,
        shell=True)

    assert "==========" in str(out)


# Command line validation tests


def test_unilist_main_too_few_args():
    with pytest.raises(subprocess.CalledProcessError):
        out = subprocess.check_output(
            "./unilist",
            stderr=subprocess.STDOUT,
            shell=True)


# Single Unicode code point argument tests

def test_unilist_main_single_valid_hexadecimal():
    out = subprocess.check_output(
        "./unilist 0020",
        stderr=subprocess.STDOUT,
        shell=True)

    assert out.decode('utf-8') == "U+0020"


def test_unilist_main_single_valid_hexadecimal_comma_delimited():
    out = subprocess.check_output(
        "./unilist -c 0020",
        stderr=subprocess.STDOUT,
        shell=True)

    assert out.decode('utf-8') == "U+0020"


def test_unilist_main_single_valid_hexadecimal_newline_delimited():
    out = subprocess.check_output(
        "./unilist -n 0020",
        stderr=subprocess.STDOUT,
        shell=True)

    assert out.decode('utf-8') == "U+0020"


def test_unilist_main_single_points_multiple_args_valid_hexadecimal():
    out = subprocess.check_output(
        "./unilist 0020 0021",
        stderr=subprocess.STDOUT,
        shell=True)

    assert out.decode('utf-8') == "U+0020 U+0021"


def test_unilist_main_single_points_multiple_args_valid_hexadecimal_comma_delimited():
    out = subprocess.check_output(
        "./unilist -c 0020 0021",
        stderr=subprocess.STDOUT,
        shell=True)

    assert out.decode('utf-8') == "U+0020,U+0021"


def test_unilist_main_single_points_multiple_args_valid_hexadecimal_newline_delimited():
    out = subprocess.check_output(
        "./unilist -n 0020 0021",
        stderr=subprocess.STDOUT,
        shell=True)

    assert out.decode('utf-8') == "U+0020\nU+0021"


def test_unilist_main_single_invalid_unicode_value():
    with pytest.raises(subprocess.CalledProcessError):
        out = subprocess.check_output(
            "./unilist ZZZZ",
            stderr=subprocess.STDOUT,
            shell=True)


def test_unilist_main_single_invalid_unicode_value_tooshort():
    with pytest.raises(subprocess.CalledProcessError):
        out = subprocess.check_output(
            "./unilist 20",
            stderr=subprocess.STDOUT,
            shell=True)


# Unicode code point range argument tests

def test_unilist_main_range_valid_hexadecimal():
    out = subprocess.check_output(
        "./unilist 0020-0022",
        stderr=subprocess.STDOUT,
        shell=True)

    assert out.decode('utf-8') == "U+0020 U+0021 U+0022"


def test_unilist_main_range_valid_hexadecimal_comma_delimited():
    out = subprocess.check_output(
        "./unilist -c 0020-0022",
        stderr=subprocess.STDOUT,
        shell=True)

    assert out.decode('utf-8') == "U+0020,U+0021,U+0022"


def test_unilist_main_range_valid_hexadecimal_newline_delimited():
    out = subprocess.check_output(
        "./unilist -n 0020-0022",
        stderr=subprocess.STDOUT,
        shell=True)

    assert out.decode('utf-8') == "U+0020\nU+0021\nU+0022"


def test_unilist_main_range_with_single_include_point():
    out = subprocess.check_output(
        "./unilist 0020-0022 0023",
        stderr=subprocess.STDOUT,
        shell=True)

    assert out.decode('utf-8') == "U+0020 U+0021 U+0022 U+0023"


def test_unilist_main_range_with_single_include_point_comma_delimited():
    out = subprocess.check_output(
        "./unilist -c 0020-0022 0023",
        stderr=subprocess.STDOUT,
        shell=True)

    assert out.decode('utf-8') == "U+0020,U+0021,U+0022,U+0023"


def test_unilist_main_range_with_single_include_point_newline_delimited():
    out = subprocess.check_output(
        "./unilist -n 0020-0022 0023",
        stderr=subprocess.STDOUT,
        shell=True)

    assert out.decode('utf-8') == "U+0020\nU+0021\nU+0022\nU+0023"


def test_unilist_main_range_invalid_start_unicode_value():
    with pytest.raises(subprocess.CalledProcessError):
        out = subprocess.check_output(
            "./unilist ZZZZ-0020",
            stderr=subprocess.STDOUT,
            shell=True)


def test_unilist_main_range_invalid_end_unicode_value():
    with pytest.raises(subprocess.CalledProcessError):
        out = subprocess.check_output(
            "./unilist 0020-ZZZZ",
            stderr=subprocess.STDOUT,
            shell=True)


def test_unilist_main_range_invalid_start_unicode_value_tooshort():
    with pytest.raises(subprocess.CalledProcessError):
        out = subprocess.check_output(
            "./unilist 20-0030",
            stderr=subprocess.STDOUT,
            shell=True)


def test_unilist_main_range_invalid_end_unicode_value_tooshort():
    with pytest.raises(subprocess.CalledProcessError):
        out = subprocess.check_output(
            "./unilist 0020-30",
            stderr=subprocess.STDOUT,
            shell=True)


# Unicode code point exclude tests

def test_unilist_main_exclude_single_value():
    out = subprocess.check_output(
        "./unilist 0020-0022 x0021",
        stderr=subprocess.STDOUT,
        shell=True)

    assert out.decode('utf-8') == "U+0020 U+0022"


def test_unilist_main_exclude_single_value_comma_delimited():
    out = subprocess.check_output(
        "./unilist -c 0020-0022 x0021",
        stderr=subprocess.STDOUT,
        shell=True)

    assert out.decode('utf-8') == "U+0020,U+0022"


def test_unilist_main_exclude_single_value_newline_delimited():
    out = subprocess.check_output(
        "./unilist -n 0020-0022 x0021",
        stderr=subprocess.STDOUT,
        shell=True)

    assert out.decode('utf-8') == "U+0020\nU+0022"


def test_unilist_main_exclude_range():
    out = subprocess.check_output(
        "./unilist 0020-0025 x0021-0024",
        stderr=subprocess.STDOUT,
        shell=True)

    assert out.decode('utf-8') == "U+0020 U+0025"


def test_unilist_main_exclude_range_comma_delimited():
    out = subprocess.check_output(
        "./unilist -c 0020-0025 x0021-0024",
        stderr=subprocess.STDOUT,
        shell=True)

    assert out.decode('utf-8') == "U+0020,U+0025"


def test_unilist_main_exclude_range_newline_delimited():
    out = subprocess.check_output(
        "./unilist -n 0020-0025 x0021-0024",
        stderr=subprocess.STDOUT,
        shell=True)

    assert out.decode('utf-8') == "U+0020\nU+0025"


def test_unilist_main_exclude_invalid_single_hexadecimal():
    with pytest.raises(subprocess.CalledProcessError):
        out = subprocess.check_output(
            "./unilist 0020-0025 xZZZZ",
            stderr=subprocess.STDOUT,
            shell=True)


def test_unilist_main_exclude_invalid_multiple_hexadecimal_start():
    with pytest.raises(subprocess.CalledProcessError):
        out = subprocess.check_output(
            "./unilist 0020-0025 xZZZZ-0030",
            stderr=subprocess.STDOUT,
            shell=True)


def test_unilist_main_exclude_invalid_multiple_hexadecimal_end():
    with pytest.raises(subprocess.CalledProcessError):
        out = subprocess.check_output(
            "./unilist 0020-0025 x0020-ZZZZ",
            stderr=subprocess.STDOUT,
            shell=True)


def test_unilist_main_exclude_invalid_single_hexadecimal_tooshort():
    with pytest.raises(subprocess.CalledProcessError):
        out = subprocess.check_output(
            "./unilist 0020-0025 x20",
            stderr=subprocess.STDOUT,
            shell=True)


def test_unilist_main_exclude_invalid_multiple_hexadecimal_start_tooshort():
    with pytest.raises(subprocess.CalledProcessError):
        out = subprocess.check_output(
            "./unilist 0020-0025 x20-0030",
            stderr=subprocess.STDOUT,
            shell=True)


def test_unilist_main_exclude_invalid_multiple_hexadecimal_end_tooshort():
    with pytest.raises(subprocess.CalledProcessError):
        out = subprocess.check_output(
            "./unilist 0020-0025 x0020-30",
            stderr=subprocess.STDOUT,
            shell=True)


